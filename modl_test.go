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
	"load":          true,
	"method":        true,
	"conditional":   true,
	"punycode":      true,
	"class":         true,
	"array":         true,
	"object_ref":    true,
	"nbArray":       true,
	"map":           true,
	"pair":          true,
	"string_method": true,
	"version":       true,
	"undefined":     true,
	"graves":        true,
	"quotes":        true,
	"escape":        true,
	"unicode":       true,
	"@keys":         true,
	"refs":          true,
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
	t.Logf("Input: %s", test.Input)
	t.Logf("Expect: %s", test.Expected)
	if err := json.Unmarshal([]byte(test.Expected), &exp); err != nil {
		t.Fatalf("Unable to parse Expected JSON: %q", err.Error())
	}
	expBits, err := json.Marshal(&exp)
	if err != nil {
		t.Fatalf("Unable to re-serialzie Expected result: %q", err)
	}
	test.Expected = string(expBits)
	if err := Unmarshal([]byte(test.Input), &got, nil); err != nil {
		t.Fatalf("Unable to unmarshal modl %q: %q", test.Input, err)
	}
	bits, err := json.Marshal(&got)
	if err != nil {
		t.Fatalf("Unable to json.Marshal parsed MODL: %q: %q", got, err)
	}
	if string(bits) != test.Expected {
		t.Fatalf("Failed Match\nExp: %s\nGot: %s", test.Expected, string(bits))
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
