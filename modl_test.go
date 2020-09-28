package modl

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

// Going with allow list given the number of failures while starting...
// will likely reverse to a block list once something starts working :cry:
var allow = map[string]bool{
	"361": true,
	"362": true,
}

// MODL features that are KNOWN to not work (right now)
var block = map[string]bool{
	"load": true,
	"method": true,
	"conditional": true,
	"punycode": true,
	"class": true,
	"array": true,
	"object_ref": true,
	"nbArray": true,
	"map": true,
	"pair": true,
	"string_method": true,
	"version": true,
	"undefined": true,
	"graves": true,
	"quotes": true,
	"escape": true,
	"unicode": true,
	"@keys": true,
	"refs": true,
}

type grammarTest struct {
	ID       string   `json:"id"`
	Input    string   `json:"input"`
	Expected string   `json:"expected_output"`
	Features []string `json:"tested_features"`
	Minimal  string   `json:"minimized_modl"`
}

func (test grammarTest) Name() string {
	if len(test.Features) == 0 {
		return test.ID
	}
	return test.ID + "_" + strings.Join(test.Features, "-")
}

func (test grammarTest) Skip() bool {
	if test.Expected == "DELETED" {
		return true
	}
	if testing.Short() && !allow[test.ID] {
		return true
	}
	for _, feat := range test.Features {
		if block[feat] {
			return true
		}
	}
	return false
}

func (test grammarTest) Test(t *testing.T) {
	var (
		array = []interface{}{}
		class = map[string]interface{}{}
		exp   = []byte(test.Expected)
	)
	t.Logf("Input: %s", test.Input)
	t.Logf("Expect: %s", test.Expected)
	if err := json.Unmarshal(exp, &array); err == nil {
		exp, err := json.Marshal(&array)
		if err != nil {
			t.Fatalf("Unable to re-serialize result: %q", err)
		}
		test.Expected = string(exp)
		test.TestArray(t)
	} else if err = json.Unmarshal(exp, &class); err == nil {
		exp, err := json.Marshal(&class)
		if err != nil {
			t.Fatalf("Unable to re-serialize result: %q", err)
		}
		test.Expected = string(exp)
		test.TestClass(t)
	} else {
		t.Fatalf("Unable to derive type of result: %q", test.Expected)
	}
}

func (test grammarTest) TestArray(t *testing.T) {
	subject := []interface{}{}
	if err := Unmarshal([]byte(test.Input), &subject, nil); err != nil {
		t.Fatalf("Unable to unmarshal modl %q: %q", test.Input, err)
	}
	bits, err := json.Marshal(&subject)
	if err != nil {
		t.Fatalf("Unable to json.Marshal response: %q: %q", subject, err)
	}
	if string(bits) != string(test.Expected) {
		t.Fatalf("Failed Match\nExp: %s\nGot: %s", string(test.Expected), string(bits))
	}
}

func (test grammarTest) TestClass(t *testing.T) {
	subject := map[string]interface{}{}
	if err := Unmarshal([]byte(test.Input), &subject, nil); err != nil {
		t.Fatalf("Unable to unmarshal modl %q: %q", test.Input, err)
	}
	bits, err := json.Marshal(&subject)
	if err != nil {
		t.Fatalf("Unable to json.Marshal response: %q: %q", subject, err)
	}
	if string(bits) != string(test.Expected) {
		t.Fatalf("Failed Match\nExp: %s\nGot: %s", string(test.Expected), string(bits))
	}
}

func TestGrammar(t *testing.T) {
	jsonTests, err := ioutil.ReadFile("grammar/tests/base_tests.json")
	if err != nil {
		t.Skip("Cannot read base_tests.json (git submodule init): " + err.Error())
	}
	var tests []grammarTest
	if err = json.Unmarshal(jsonTests, &tests); err != nil {
		t.Fatal("Unable to deserialize tests: " + err.Error())
	}
	for _, test := range tests {
		if !test.Skip() {
			t.Run(test.Name(), test.Test)
		}
	}
}
