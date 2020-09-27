package modl

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

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
	t.Logf("Parsing %q", test.Input)
}

func TestGrammar(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	jsonTests, err := ioutil.ReadFile("grammar/tests/base_tests.json")
	if err != nil {
		t.Skip("Cannot read base_tests.json (git submodule init): " + err.Error())
	}
	var tests []grammarTest
	if err = json.Unmarshal(jsonTests, &tests); err != nil {
		t.Fatal("Unable to deserialize tests: " + err.Error())
	}
	for _, test := range tests {
		t.Run(test.Name(), test.Test)
	}
}
