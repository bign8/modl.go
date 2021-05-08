package unpacker

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	if state.err != nil {
		return nil, state.err
	}
	return json.Marshal(v)
}

type unpackState struct {
	Idx []interface{} `json:"?"`
	ctx Unpacker
	dec *json.Decoder
	err error
}

func (state unpackState) value() interface{} {
	if state.err != nil {
		return nil
	}
	token, err := state.dec.Token()
	if err != nil {
		state.err = err
		return nil
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
		return state.string(v)
	case float64:
		// fmt.Printf("Got an int value: %f\n", token)
		return v
	default:
		panic(fmt.Sprintf("Unknown type: %T", token))
	}
}

func (state unpackState) object() map[string]interface{} {
	if state.err != nil {
		return nil
	}

	// TODO: skip over `?` in top level object (maybe always for now :badpokerface:)
	o := map[string]interface{}{}
	for state.dec.More() {
		k := state.value().(string)
		v := state.value()

		o[k] = v

		// TODO: apply transforms to value
	}
	state.delim('}')
	return o
}

func (state unpackState) array() []interface{} {
	if state.err != nil {
		return nil
	}
	o := []interface{}{}
	for state.dec.More() {
		o = append(o, state.value())
	}
	state.delim(']')
	return o
}

func (state unpackState) string(in string) string {
	if state.err != nil {
		return ""
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
