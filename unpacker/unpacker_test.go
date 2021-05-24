package unpacker

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"
)

type UnpackerTest struct {
	ID     string
	Input  json.RawMessage
	Subs   json.RawMessage
	Trans  json.RawMessage
	Output json.RawMessage
	Skip   bool
}

func (test *UnpackerTest) Run(t *testing.T) {
	if testing.Short() && test.Skip {
		t.Skip("manually disabled test")
	}
	u := &Unpacker{}
	if len(test.Trans) != 0 {
		trans, err := ParseTransforms([]byte(test.Trans))
		if err != nil {
			t.Fatalf("Unable to parse transforms: %v", err)
		}
		u.Transforms = trans
		// fmt.Printf("Got Transforms: %v\n", trans)
	}
	if len(test.Subs) != 0 {
		subs := make(Substitution)
		if err := json.Unmarshal([]byte(test.Subs), &subs); err != nil {
			t.Fatalf("Unable to parse substition: %v", err)
		}
		u.AddSubs(subs)
	}
	output, err := u.Unpack([]byte(test.Input))
	if err != nil {
		t.Fatalf("Unable to unpack: %v", err)
	}
	exp := test.out()
	if !bytes.Equal(output, exp) {
		t.Fatalf("Unexpected output:\n\tWant: %s\n\tGot:  %s", string(exp), string(output))
	}
}

func (test UnpackerTest) out() []byte {
	var obj interface{}
	err := json.Unmarshal([]byte(test.Output), &obj)
	if err != nil {
		panic(err)
	}
	bits, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return bits
}

var skip = map[string]bool{
	"Unpacker-07": true, // transform with variadic-index reference
	"Unpacker-13": true, // arrayItems + nested
	"Unpacker-14": true, // arrayItems + nested
	"Unpacker-":   true,
}

func TestUnpacker(t *testing.T) {
	jsonTests, err := ioutil.ReadFile("tests.json")
	if err != nil {
		t.Skip("Cannot read base_tests.json (git submodule init): " + err.Error())
	}
	var tests []UnpackerTest
	if err = json.Unmarshal(jsonTests, &tests); err != nil {
		t.Fatal("Unable to deserialize tests: " + err.Error())
	}
	for _, test := range tests {
		test.Skip = skip[test.ID]
		t.Run(test.ID, test.Run)
	}
}
