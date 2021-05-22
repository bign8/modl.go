package unpacker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Unpacker struct {
	Transforms map[string]Transform
	Subs       Substitution
}

type Transform struct {
	Assign  []string               `json:"assignKeys,omitempty"`
	Items   string                 `json:"arrayItems,omitempty"`
	Key     string                 `json:"rewriteKey,omitempty"`
	Return  map[string]interface{} `json:"replacePair,omitempty"`  // can be "null" to nuke a pair (how?)
	Rewrite interface{}            `json:"rewriteValue,omitempty"` // freeform object
}

type Substitution map[string]interface{}

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

func (u Unpacker) Unpack(source []byte) ([]byte, error) {
	state := &unpackState{ctx: u}

	// pass one, extract the question mark lookup index
	if err := json.Unmarshal(source, state); err != nil {
		return nil, err
	}
	// TODO: pre-process the question mark values

	// pass two, process the data
	state.dec = json.NewDecoder(bytes.NewReader(source))
	v := state.value()
	// TODO: trap errors?
	return json.Marshal(v)
}

type unpackState struct {
	Idx []interface{} `json:"?"`
	ctx Unpacker
	dec *json.Decoder
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
		return state.string(v, state.ctx.Subs)
	case float64:
		// fmt.Printf("Got an int value: %f\n", token)
		return v
	default:
		panic(fmt.Sprintf("Unknown type: %T", token))
	}
}

func (state unpackState) object() map[string]interface{} {
	o := map[string]interface{}{}
	for state.dec.More() {
		k := state.value().(string)
		v := state.value()
		if k == "?" {
			continue
		}
		state.transform(o, k, v)
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
	for i, v := range state.Idx {
		key := "%" + strconv.Itoa(i)
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

	// TODO: dot lookups
	// TODO: closing % replacements
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
	trans, ok := state.ctx.Transforms[key]
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
		for i, v := range list {
			n := map[string]interface{}{} // TODO: preallocate size
			state.transform(n, trans.Items, v)
			list[i] = n
		}
		dest[key] = list
		return // not much else we can do here
	}

	// 1. Assign(assignKeys) converts arrays to objects
	if len(trans.Assign) != 0 {
		list, ok := value.([]interface{})
		if !ok {
			fmt.Print("assignKeys for " + key + " but didn't get a list")
			return
		}
		if len(list) != len(trans.Assign) {
			panic(fmt.Sprintf("assignKeys for %s expected %d but got %d keys", key, len(trans.Assign), len(list)))
		}
		newValue := make(map[string]interface{}, len(list))
		for i, v := range list {
			state.transform(newValue, trans.Assign[i], v) // recurse to assign key/values
		}
		value = newValue
	}

	// 2: replacePair and exit
	if trans.Return != nil {
		ctx := map[string]interface{}{
			"self": value,
		}
		if m, ok := value.(map[string]interface{}); ok {
			for k, v := range m {
				ctx[k] = v
			}
		}
		// fmt.Printf("Got replacer: %v\n", value)
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
	if trans.Rewrite != nil {
		ctx := map[string]interface{}{
			"self": value,
		}
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
