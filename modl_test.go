package modl

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

// base_tests.json that are currently failing
var skip = map[string]bool{
	"b216": true, // null
	"e058": true, // multiple map keys
}

type grammarTest struct {
	ID       string   `json:"id"`
	Input    string   `json:"input"`
	Expected string   `json:"expected_output"`
	Features []string `json:"tested_features"`
	Minimal  string   `json:"minimized_modl"`
}

func (test grammarTest) Test(t *testing.T) {
	if testing.Short() && skip[test.ID] {
		t.Skip("excluded in skip array")
	}
	var (
		exp interface{}
		got interface{}
	)
	if test.Expected[0] == '[' {
		exp = []interface{}{}
		got = []interface{}{}
	} else {
		exp = map[string]interface{}{}
		got = map[string]interface{}{}
	}
	if err := json.Unmarshal([]byte(test.Expected), &exp); err != nil {
		t.Fatalf("Unable to parse Expected JSON: %q", err.Error())
	}
	expBits, err := json.Marshal(&exp)
	if err != nil {
		t.Fatalf("Unable to re-serialzie Expected result: %q", err)
	}
	test.Expected = string(expBits)
	if err := Unmarshal([]byte(test.Input), &got); err != nil {
		t.Fatalf("Unable to unmarshal modl %q: %q", test.Input, err)
	}
	bits, err := json.Marshal(&got)
	if err != nil {
		t.Fatalf("Unable to json.Marshal parsed MODL: %q: %q", got, err)
	}
	if string(bits) != test.Expected {
		t.Fatalf("Failed Match\nInp: %s\nExp: %s\nGot: %s", test.Input, test.Expected, string(bits))
	} else {
		t.Logf("Success!\nInp: %s\nGot: %s", test.Input, string(bits))
	}
}

func TestGrammarBase(t *testing.T) {
	jsonTests, err := ioutil.ReadFile("grammar/tests/base_tests.json")
	if err != nil {
		t.Skip("Cannot read base_tests.json (git submodule init): " + err.Error())
	}
	var tests []grammarTest
	if err = json.Unmarshal(jsonTests, &tests); err != nil {
		t.Fatal("Unable to deserialize tests: " + err.Error())
	}
	for _, test := range tests {
		for len(test.ID) < 3 {
			test.ID = "0" + test.ID
		}
		test.ID = "b" + test.ID
		t.Run(test.ID, test.Test)
	}
}

func TestGrammarExtra(t *testing.T) {
	jsonTests, err := ioutil.ReadFile("grammar/tests/extra_tests.json")
	if err != nil {
		t.Skip("Cannot read extra_tests.json (git submodule init): " + err.Error())
	}
	var tests []grammarTest
	if err = json.Unmarshal(jsonTests, &tests); err != nil {
		t.Fatal("Unable to deserialize tests: " + err.Error())
	}
	for _, test := range tests {
		for len(test.ID) < 3 {
			test.ID = "0" + test.ID
		}
		test.ID = "e" + test.ID
		t.Run(test.ID, test.Test)
	}
}

func chk(tb testing.TB, x interface{}) {
	y := reflect.ValueOf(x)
	tb.Log(y.String())
}

func TestBadInterfaceReflection(t *testing.T) {
	// TODO: refactor tests so this isn't a problem
	x := map[string]interface{}{}
	var y interface{}
	y = x
	chk(t, &y)
	if err := json.Unmarshal([]byte("{\"a\": \"A\"}"), &y); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	t.Logf("Something: %#v", x)
}
