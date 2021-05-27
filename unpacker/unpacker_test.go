package unpacker

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"
)

type UnpackerTest struct {
	ID        string
	Input     json.RawMessage
	Subs      Substitution
	Trans     map[string]Transform
	TransFile string `json:"trans_file"`
	Output    json.RawMessage
	Skip      bool
}

func (test *UnpackerTest) Run(t *testing.T) {
	must := func(err error) {
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
	if testing.Short() && test.Skip {
		t.Skip("manually disabled test")
	}
	u := &Unpacker{}
	if len(test.TransFile) != 0 {
		bits, err := ioutil.ReadFile(test.TransFile)
		must(err)
		test.Trans, err = ParseTransforms(bits)
		must(err)
	}
	if len(test.Trans) != 0 {
		u.Transforms = test.Trans
		// fmt.Printf("Got Transforms: %v\n", u.Transforms)
	}
	if len(test.Subs) != 0 {
		u.AddSubs(test.Subs)
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
	"Supplementary-2": true, // % escapes? (stronger string parsing)
	"Supplementary-3": true, // unresolved refs (stronger string parsing)
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
