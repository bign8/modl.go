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
	"DELETED":       true,
	"load":          true,
	"object_ref":    true,
	"conditional":   true,
	"class":         true,
	"method":        true,
	"refs":          true,
	"string_method": true,
	"punycode":      true, // needs methods -> will use golang.org/x/net/idna::ToUnicode
}

// Tests on un-supported features, that pass due to noops
var allow = map[string]bool{
	"003": true, // noop load
	"042": true, // noop class
	"085": true, // noop class
	"178": true, // noop load
	"179": true, // noop load
	"270": true, // noop load
	"271": true, // noop load
	"278": true, // noop method
	"333": true, // noop load
	"343": true, // noop replace
}

// Tests with supported features that don't quite work right :cry:
var skip = map[string]bool {
	"167": true, // missing_label => object_ref/conditional
	"211": true, // array with null => {"o":[1,2,null,4,5]}
	"212": true, // array with null => {"o":[1,2,null,4,5]}
	"236": true, // missing_label => object_ref
	"246": true, // complicated array => {"x":[1,[2,3,4]]}
	"248": true, // complicated array + null => x=[1;;;;;;;2:::::::::3;;;;;;;;;;4:::::::::5]
	"283": true, // missing_label ref
	"284": true, // missing_label ref
	"314": true, // missing_label ref
	"315": true, // missing_label ref
	"323": true, // double parsed escape sequence on unicode (emailing with MODL maintainers)
	"341": true, // complicated array => [[1,2,3],{"b":"c"},[4,5,6],{"d":"f"}]
	"363": true, // unlabled refs
	"364": true, // unlabled class
	"365": true, // unlabled refs
	"366": true, // unlabled refs
	"367": true, // unlabled refs
	"368": true, // unlabled refs
	"369": true, // unlabled refs
	"370": true, // unlabled class
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