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
}

// Tests on un-supported features, that pass due to noops
var allow = map[string]bool{
	"3":   true, // noop load
	"42":  true, // noop class
	"85":  true, // noop class
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
	"22":  true, // punycode (not 100% working yet)
	"61":  true, // complicated map
	"63":  true, // complicated map
	"64":  true, // complicated map
	"72":  true, // version (quoted key?)
	"73":  true, // version (quoted key?)
	"91":  true, // complicated objects + array
	"118": true, // array of arrays
	"167": true, // "undefined", object_ref/conditional
	"187": true, // array of objects
	"188": true, // array of objects
	"192": true, // array of objects
	"211": true, // array with null
	"212": true, // array with null
	"236": true, // "undefined", object_ref
	"246": true, // complicated array
	"248": true, // complicated array + null
	"283": true, // "graves" with ref
	"284": true, // "quotes" with ref
	"298": true, // weird mixed quotes + graves
	"314": true, // "escape" with ref
	"315": true, // "escape" with ref
	"323": true, // double parsed escape sequence on unicode
	"341": true, // complicated array
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