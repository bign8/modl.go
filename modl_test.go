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

func (test grammarTest) Test(t *testing.T) {
	var (
		array = []interface{}{}
		class = map[string]interface{}{}
		exp   = []byte(test.Expected)
	)
	if err := json.Unmarshal(exp, &array); err == nil {
		exp, err := json.Marshal(array)
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
		t.Fatalf("Failed Match\nExp: %q\nGot: %q", string(test.Expected), string(bits))
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
		t.Fatalf("Failed Match\nExp: %q\nGot: %q", string(test.Expected), string(bits))
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
		if test.Expected == "DELETED" {
			continue
		}
		if testing.Short() && !allow[test.ID] {
			continue
		}
		t.Run(test.Name(), test.Test)
	}
}
