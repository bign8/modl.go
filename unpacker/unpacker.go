package unpacker

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Unpacker struct {
	Transforms []Transform
	Subs       Substitution
}

type Transform struct {
	Assign  []string               `json:"assignKeys,omitempty"`
	Items   string                 `json:"arrayItems,omitempty"`
	Key     string                 `json:"rewriteKey,omitempty"`
	Name    string                 `json:"-"`                      // the key of the json object
	Return  map[string]interface{} `json:"replacePair,omitempty"`  // can be "null" to nuke a pair (how?)
	Rewrite interface{}            `json:"rewriteValue,omitempty"` // freeform object
}

type Substitution map[string]interface{}

func ParseTransforms(transform []byte) ([]Transform, error) {
	v := map[string]Transform{}
	if err := json.Unmarshal(transform, &v); err != nil {
		return nil, err
	}
	t := make([]Transform, 0, len(v))
	for key, value := range v {
		value.Name = key
		t = append(t, value)
	}
	return t, nil
}

func (u *Unpacker) AddSubs(subs Substitution) {
	if u.Subs == nil {
		u.Subs = make(Substitution)
	}
	for key, value := range subs {
		u.Subs[key] = value
	}
}

func (u *Unpacker) AddTransform(trans ...Transform) {
	u.Transforms = append(u.Transforms, trans...)
}

func (u Unpacker) Unpack(source []byte) ([]byte, error) {
	obj := map[string]interface{}{}
	if err := json.Unmarshal(source, &obj); err != nil {
		return nil, err
	}
	index, hasIndex := obj["?"]
	delete(obj, "?")
	fmt.Printf("Has index: %t: %v\n", hasIndex, index)

	// note the special ? key in the source object (which may or may not be present)
	return nil, errors.New("TODO")
}
