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
	"3":   true,
	"62":  true,
	"79":  true,
	"132": true,
	"144": true,
	"146": true,
	"178": true,
	"179": true,
	"190": true,
	"191": true,
	"193": true,
	"206": true,
	"278": true,
	"286": true,
	"297": true,
	"316": true,
	"321": true,
	"322": true,
	"324": true,
	"325": true,
	"332": true,
	"333": true,
	"334": true,
	"343": true,
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

func (test grammarTest) Skip() bool {
	if test.Expected == "DELETED" {
		return true
	}
	if testing.Short() {
		return !allow[test.ID]
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
		t.Fatalf("Failed Match\nInp: %s\nExp: %s\nGot: %s", test.Input, test.Expected, string(bits))
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
