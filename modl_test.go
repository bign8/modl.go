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
	"162": true, // noop ref
	"178": true, // noop load
	"179": true, // noop load
	"270": true, // noop load
	"271": true, // noop load
	"278": true, // noop method
	"333": true, // noop load
	"343": true, // noop replace
	"349": true, // noop ref
	"351": true, // noop ref
	"353": true, // noop ref
	"355": true, // noop ref
}

// Tests with supported features that don't quite work right :cry:
// WHERE question == "object index" (5.2) of modl spec
var skip = map[string]bool{
	"040": true, // WIP: obj-ref - _test=[[a;b]];letters2=%test.0.0 - dots, array
	"048": true, // WIP: obj-ref - ?[[a;b;c]];letters=%0 - question
	"049": true, // WIP: obj-ref - ?=[a;b;c;d]:[1;2;3;4;5];test=%1.0 - question, dots
	"050": true, // WIP: obj-ref - _test=123;object(print_test = %test.test) - dots, stop, early
	"051": true, // WIP: obj-ref - <OMITTED> - question
	"058": true, // WIP: obj-ref - _test=(a=b=c=d=f);x=%test.a.b.c.d - dots
	"059": true, // WIP: obj-ref - a(b(c(d(e(f=1)))));testing=%a.b.c.d.e.f - dots
	"060": true, // WIP: obj-ref - _test=(a=b=c=d=f);testing=%test.a.b.c.d - dots
	"096": true, // WIP: obj-ref - _person(name(first=John;last=Smith));say=%person.name.first - dots
	"098": true, // WIP: obj-ref - _C=gb;_COUNTRIES(gb=ASDF);name=%COUNTRIES.%C - dots, double lookup
	"100": true, // WIP: obj-ref - _person(name(first="John"));a=%person.name.first - dots
	"101": true, // WIP: obj-ref - ?=[a;b;c;d]:[1;2;3;4;5];test=%1 - question
	"114": true, // WIP: obj-ref - ?[[a;b;c];[one;two;three]];letters=%0;numbers=%1 - question
	"115": true, // WIP: obj-ref - ?=[a;b;c]:[one;two;three];letters=%0;numbers=%1 - question
	"116": true, // WIP: obj-ref - ?[a;b;c];letters=%0 - question
	"120": true, // WIP: obj-ref - <OMITTED> - question
	"124": true, // WIP: obj-ref - _test[a;b;c];alex=%test.0 - dots, array
	"125": true, // WIP: obj-ref - ?[a;b;c];alex=%0 - question
	"140": true, // WIP: obj-ref - question
	"141": true, // WIP: obj-ref - question
	"143": true, // WIP: obj-ref - question
	"161": true, // WIP: obj-ref - flat-return-null?
	"167": true, // missing_label => object_ref/conditional
	"217": true, // WIP: obj-ref - dots, array
	"218": true, // WIP: obj-ref - dots
	"221": true, // WIP: obj-ref - dots, array, object
	"222": true, // WIP: obj-ref - dots, array, object
	"223": true, // WIP: obj-ref - dots, object
	"236": true, // missing_label => object_ref
	"237": true, // WIP: obj-ref - dots
	"247": true, // WIP: obj-ref - dots, array
	"273": true, // WIP: obj-ref - string functions
	"283": true, // missing_label ref
	"284": true, // missing_label ref
	"285": true, // WIP: obj-ref - string functions
	"299": true, // WIP: obj-ref - string functions
	"308": true, // WIP: obj-ref - weird, closure [SHOULD WORK?] - number, stop on space
	"312": true, // WIP: obj-ref - weird, closure [SHOULD WORK?] - number, stop on grave
	"323": true, // double parsed escape sequence on unicode (emailing with MODL maintainers)
	"329": true, // WIP: obj-ref - load, num-record, question
	"330": true, // WIP: obj-ref - dots, array
	"331": true, // WIP: obj-ref - load
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
