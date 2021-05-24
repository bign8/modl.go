// Package unpacker implements the logic described at https://www.unpacker.uk/ to unpack compressed JSON objects.
package unpacker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Unpakcker holds a set of Transforms and Substitutions needed to unpack compressed JSON.
// This structure should allow the the Transform and Substitution object to only need to be parsed once,
// while allowing multiple unpackings using the same configuration.
type Unpacker struct {
	Transforms map[string]Transform
	Subs       Substitution
}

// A Transform holds information on how to modify various JSON objects.
//
// docs: https://www.unpacker.uk/specification#transformation-object
//
// Note: field tags are included on this object for convience as the real work is performed in Un/MarshalJSON.
type Transform struct {
	Assign      []string               `json:"assignKeys,omitempty"`
	Items       string                 `json:"arrayItems,omitempty"`
	Key         string                 `json:"rewriteKey,omitempty"`
	Return      map[string]interface{} `json:"replacePair,omitempty"` // can be "null" to nuke a pair (how?)
	ReturnNull  bool                   `json:"-"`
	ReturnStr   string                 `json:"-"`                      // replacePair: "%self"
	Rewrite     interface{}            `json:"rewriteValue,omitempty"` // freeform object
	RewriteNull bool                   `json:"-"`
	Nesting     map[string]Transform   `json:"-"`
}

func (t *Transform) UnmarshalJSON(bits []byte) error {
	// TODO: emit errors if the types are wrong
	trans := make(map[string]interface{}) // todo: map string to json.RawMessage?
	if err := json.Unmarshal(bits, &trans); err != nil {
		return err
	}
	if assign, ok := trans["assignKeys"].([]interface{}); ok {
		t.Assign = make([]string, 0, len(assign))
		for _, a := range assign {
			t.Assign = append(t.Assign, a.(string))
		}
		delete(trans, "assignKeys")
	}
	if items, ok := trans["arrayItems"].(string); ok {
		t.Items = items
		delete(trans, "arrayItems")
	}
	if key, ok := trans["rewriteKey"].(string); ok {
		t.Key = key
		delete(trans, "rewriteKey")
	}
	if replace, ok := trans["replacePair"]; ok {
		if replace == nil {
			t.ReturnNull = true
		} else if str, ok := replace.(string); ok {
			t.ReturnStr = str
		} else if rep, ok := replace.(map[string]interface{}); ok {
			t.Return = rep
		}
		delete(trans, "replacePair")
	}
	if rewrite, ok := trans["rewriteValue"]; ok {
		if rewrite == nil {
			t.RewriteNull = true
		} else {
			t.Rewrite = rewrite
		}
		delete(trans, "rewriteValue")
	}
	if len(trans) > 0 {
		t.Nesting = make(map[string]Transform)
		bits, err := json.Marshal(trans)
		if err != nil {
			return err
		}
		return json.Unmarshal(bits, &t.Nesting)
	}
	return nil
}

func (t Transform) MarshalJSON() ([]byte, error) {
	trans := make(map[string]interface{})
	if len(t.Assign) > 0 {
		trans["assignKeys"] = t.Assign
	}
	if t.Items != "" {
		trans["arrayItems"] = t.Items
	}
	if t.Key != "" {
		trans["rewriteKey"] = t.Key
	}
	if t.ReturnNull {
		trans["replacePair"] = nil
	} else if t.ReturnStr != "" {
		trans["replacePair"] = t.ReturnStr
	} else if t.Return != nil {
		trans["replacePair"] = t.Return
	}
	if t.RewriteNull {
		trans["rewriteValue"] = nil
	} else if t.Rewrite != nil {
		trans["rewriteValue"] = t.Rewrite
	}
	for k, v := range t.Nesting {
		trans[k] = v
	}
	return json.Marshal(trans)
}

func (t Transform) String() string {
	x, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}
	return string(x)
}

// A Substitution object provides a method of removing repeditive data by having shortened keys.
//
// docs: https://www.unpacker.uk/specification#substitution-object
type Substitution map[string]interface{}

// ParseTransforms parses a set of json tranforms.
// Transforms can be embeded, so the bulk of this logic is to properly parse embedded transforms.
// Example: {"k": {"e": {"y": {/*transform object*/}}}}
func ParseTransforms(transforms []byte) (map[string]Transform, error) {
	v := map[string]Transform{}
	return v, json.Unmarshal(transforms, &v)
}

func (u *Unpacker) AddSubs(subs Substitution) {
	if u.Subs == nil {
		u.Subs = make(Substitution)
	}
	for key, value := range subs {
		u.Subs[key] = value
	}
}

func (u *Unpacker) AddTransform(key string, trans Transform) {
	// TODO: create map if not exists
	// TODO: panic if key already exists
	u.Transforms[key] = trans
}

// Unpack expands JSON using the transforms and substitutions defined on the unpacker object.
// Additionally, a viariadic index in the source data is also used to reduce size.
//
// docs: https://www.unpacker.uk/specification
func (u Unpacker) Unpack(source []byte) ([]byte, error) {
	state := &unpackState{
		fullMem:   make(map[string]interface{}),
		noVariMem: make(map[string]interface{}),
	}
	state.trans = append(state.trans, u.Transforms)

	// pass one, extract the variable-index and memoize everything for `/compact.` namespaced vars
	var compact interface{}
	if err := json.Unmarshal(source, &compact); err != nil {
		return nil, err
	}

	// memoize string replacements
	if obj, ok := compact.(map[string]interface{}); ok {
		augment("", obj["?"], state.fullMem)
	}
	augment("", u.Subs, state.fullMem)
	augment("", u.Subs, state.noVariMem)
	augment("/subs.", u.Subs, state.fullMem)
	augment("/subs.", u.Subs, state.noVariMem)
	augment("/compact.", compact, state.fullMem)
	augment("/compact.", compact, state.noVariMem)
	// fmt.Printf("Values: %#v\n", state.mem)

	// pass two, process the data
	state.dec = json.NewDecoder(bytes.NewReader(source))
	v := state.value()
	// TODO: trap errors?
	return json.Marshal(v)
}

type unpackState struct {
	dec       *json.Decoder
	fullMem   map[string]interface{}
	noVariMem map[string]interface{}
	trans     []map[string]Transform
}

func (state unpackState) value() interface{} {
	token, err := state.dec.Token()
	if err != nil {
		panic(err)
	}
	switch v := token.(type) {
	case json.Delim:
		// fmt.Printf("Got a delimiter: %v\n", token)
		switch v.String() {
		case "{":
			return state.object()
		case "[":
			return state.array()
		default:
			panic("Unknown delim string: " + v.String())
		}
	case string:
		// fmt.Printf("Got a string value: %q\n", token)
		return state.string(v, state.fullMem)
	case float64:
		// fmt.Printf("Got an int value: %f\n", token)
		return v
	case nil:
		return nil
	default:
		panic(fmt.Sprintf("Unknown type: %T", token))
	}
}

func (state unpackState) object() map[string]interface{} {
	o := map[string]interface{}{}
	for state.dec.More() {
		k := state.value().(string)

		// Push transform state (for use when processing value)
		var has_arr_trans bool
		trans, has_trans := state.getTransform(k)
		if has_trans {
			state.trans = append(state.trans, trans.Nesting)
		}
		if has_trans && trans.Items != "" {
			var arr_trans Transform
			arr_trans, has_arr_trans = state.getTransform(trans.Items)
			if has_arr_trans {
				state.trans = append(state.trans, arr_trans.Nesting)
			}
		}

		// Parse the value!
		v := state.value()

		// Pop arr_trans
		if has_arr_trans {
			state.trans = state.trans[:len(state.trans)-1]
		}

		// if we aren't the variadic index, transform data (time saver?)
		if k != "?" {
			state.transform(o, k, v)
		}

		// Pop transform state
		if has_trans {
			state.trans = state.trans[:len(state.trans)-1]
		}
	}
	state.delim('}')
	return o
}

func (state unpackState) array() []interface{} {
	o := []interface{}{}
	for state.dec.More() {
		o = append(o, state.value())
	}
	state.delim(']')
	return o
}

func (state unpackState) string(in string, subz map[string]interface{}) interface{} {
	// NOTE: state is not used, but is left as a potential memory optimization
	// (array of contexts or logic to switch between allowing viariadic or not)
	for k, v := range subz {
		key := "%" + k
		if in == key || in == key+"%" {
			return v
		}
		if s, ok := v.(string); ok {
			if strings.HasSuffix(in, key) {
				in = strings.TrimSuffix(in, key) + s
			}
			if strings.Contains(in, key+" ") {
				in = strings.Replace(in, key+" ", s+" ", -1)
			}
			if strings.Contains(in, key+"%") {
				in = strings.Replace(in, key+"%", s, -1)
			}
		}
	}
	// assume some part of the full thing was a key (likely a bug, need a beter way to do this)
	if strings.HasPrefix(in, "%") {
		return nil
	}
	return in
}

func (state unpackState) delim(delim json.Delim) {
	cl, err := state.dec.Token()
	if err != nil {
		panic(err)
	}
	if cl != delim {
		panic(fmt.Sprintf("Expected %v; got %v", delim, cl))
	}
}

func (state unpackState) transform(dest map[string]interface{}, key string, value interface{}) {
	trans, ok := state.getTransform(key)
	if !ok {
		dest[key] = value
		return
	}

	// 0. Items(arrayItems) convert array of arrays to an array of objects
	if trans.Items != "" {
		list, ok := value.([]interface{})
		if !ok {
			panic(fmt.Sprintf("arrayItems set need an array; %T", value))
		}

		// Push transform state (for use when processing value)
		sub_trans, has_trans := state.getTransform(trans.Items)
		if has_trans {
			state.trans = append(state.trans, map[string]Transform{
				trans.Items: sub_trans,
			})
		}

		// todo: push context of trans.Items onto stack@!
		for i, v := range list {
			n := map[string]interface{}{} // TODO: preallocate size
			state.transform(n, trans.Items, v)
			list[i] = n
		}

		// Pop transform state
		if has_trans {
			state.trans = state.trans[:len(state.trans)-1]
		}

		dest[key] = list
		return // not much else we can do here
	}

	// 0.5: Start building context (assign keys can modify some default values here)
	ctx := make(map[string]interface{}, len(state.fullMem))
	for k, v := range state.noVariMem {
		ctx[k] = v
	}

	// 1. Assign(assignKeys) converts arrays to objects
	if list, is_list := value.([]interface{}); len(trans.Assign) != 0 && is_list {
		// iff value is an array, convert the keys as specified
		if len(list) != len(trans.Assign) {
			panic(fmt.Sprintf("assignKeys for %s expected %d but got %d keys", key, len(trans.Assign), len(list)))
		}
		newValue := make(map[string]interface{}, len(list))
		for i, v := range list {
			state.transform(newValue, trans.Assign[i], v) // recurse to assign key/values
		}
		value = newValue
	} else if m, is_map := value.(map[string]interface{}); len(trans.Assign) != 0 && is_map {
		// iff the object is a map, make sure missing keys are assigned as null (ctx resolvers use this later)
		for _, k := range trans.Assign {
			if _, ok := m[k]; !ok {
				m[k] = nil
			}
		}
	}

	// 1.5: update ctx as necessary
	ctx["self"] = value
	augment("", value, ctx)
	// fmt.Printf("Got context: %v\n", ctx)

	// 2: replacePair and exit
	if trans.ReturnNull {
		return // if replacePair is null, the key should not be assigned
	} else if trans.ReturnStr != "" {
		obj := state.string(trans.ReturnStr, ctx)
		if m, ok := obj.(map[string]interface{}); ok {
			for k, v := range m {
				dest[k] = v
			}
			return
		}
		panic(fmt.Sprintf("Got unexpected ReturnString type: %T", obj))
	} else if trans.Return != nil {
		for k, v := range trans.Return {
			// TODO: smarter string resolves
			if s, ok := v.(string); ok {
				v = state.string(s, ctx)
			}
			dest[k] = v
		}
		return
	}

	// 3: rewriteValue
	if trans.RewriteNull {
		dest[key] = nil // if rewriteValue is null, assign value to null
		return
	} else if trans.Rewrite != nil {
		if m, ok := trans.Rewrite.(map[string]interface{}); ok {
			n := map[string]interface{}{} // ensure we don't duplicate the object
			for k, v := range m {
				if s, ok := v.(string); ok {
					v = state.string(s, ctx)
				}
				n[k] = v
			}
			value = n
		} else if s, ok := trans.Rewrite.(string); ok {
			value = state.string(s, ctx)
		} else {
			panic(fmt.Sprintf("got an unexpected rewriteValue type: %T", trans.Rewrite))
		}
	}

	// 4: rewriteKey, replace key if need be
	if trans.Key != "" {
		key = trans.Key
	}

	// 5: actually assign object
	dest[key] = value
}

func augment(prefix string, thing interface{}, target map[string]interface{}) {
	switch value := thing.(type) {
	case []interface{}:
		for i, v := range value {
			x := prefix + strconv.Itoa(i)
			target[x] = v
			augment(x+".", v, target)
		}
	case map[string]interface{}:
		for k, v := range value {
			x := prefix + k
			target[x] = v
			augment(x+".", v, target)
		}
	case Substitution: // map[string]interface{}
		for k, v := range value {
			x := prefix + k
			target[x] = v
			augment(x+".", v, target)
		}
	}
}

func (state unpackState) getTransform(key string) (Transform, bool) {
	for i := len(state.trans) - 1; i > 0; i-- {
		if t, ok := state.trans[i][key]; ok {
			return t, ok
		}
	}
	t, ok := state.trans[0][key]
	return t, ok
}
