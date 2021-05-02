package modl

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

// Known unsupported features; tests with this label are skipped
var unsupported = map[string]bool{
	// "DELETED": true,
	// "load":          true,
	// "conditional":   true,
	// "class":         true,
	// "method":        true,
	// "string_method": true,
	// "punycode":      true, // needs methods -> will use golang.org/x/net/idna::ToUnicode
}

// Tests on un-supported features, that pass due to noops
var allow = map[string]bool{
	// "003": true, // noop load
	// "042": true, // noop class
	// "085": true, // noop class
	// "129": true, // additional-label: conditional
	// "130": true, // additional-label: conditional
	// "178": true, // noop load
	// "179": true, // noop load
	// "270": true, // noop load
	// "271": true, // noop load
	// "278": true, // noop method
	// "333": true, // noop load
}

// Tests with supported features that don't quite work right :cry:
// WHERE question == "object index" (5.2) of modl spec
var skip = map[string]bool{
	"006": true, // pair
	"007": true, // pair
	"008": true, // pair
	"013": true, // pair
	"020": true, // pair
	"026": true, // pair
	"028": true, // pair
	"029": true,
	"031": true,
	"032": true,
	"042": true,
	"044": true,
	"049": true,
	"056": true,
	"057": true,
	"058": true,
	"059": true,
	"060": true,
	"063": true,
	"064": true,
	"070": true,
	"084": true,
	"086": true,
	"087": true,
	"091": true,
	"095": true,
	"096": true,
	"103": true,
	"105": true,
	"106": true,
	"110": true,
	"111": true,
	"112": true,
	"115": true,
	"120": true,
	"123": true,
	"124": true,
	"127": true,
	"129": true,
	"130": true,
	"131": true,
	"136": true,
	"139": true,
	"149": true,
	"150": true,
	"157": true,
	"188": true,
	"189": true,
	"191": true,
	"197": true,
	"198": true,
	"199": true,
	"200": true,
	"201": true,
	"204": true,
	"205": true, // pair
	"206": true, // pair
	"207": true, // pair
	"216": true, // null
	"219": true, // pair
	"225": true,
	"226": true, // tilde escape
	"227": true, // tilde escape
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
		for _, feat := range test.Features {
			if unsupported[feat] {
				return !allow[test.ID]
			}
		}
		return skip[test.ID]
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
	} else {
		t.Logf("Success!\nInp: %s\nGot: %s", test.Input, string(bits))
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
		for len(test.ID) < 3 {
			test.ID = "0" + test.ID
		}
		if !test.Skip() {
			t.Run(test.Name(), test.Test)
		}
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
