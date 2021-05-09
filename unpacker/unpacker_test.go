package unpacker

import (
	"bytes"
	"encoding/json"
	"testing"
)

var unpackerTests = []UnpackerTest{
	{
		Name: "all",
		Skip: true,
		Input: `{
			"?": ["https://www.widgetcompany.com", "team", "janesmith", "johnwilson", "dashnaanand"],
			"o": [
			  "Widget Company Ltd",
			  "Making the best widgets",
			  "%0",
			  [
				["Jane Smith", "%ceo", "%0%/%1%/%2", "%2", "%2%widgets"],
				["John Wilson", "%cto", "%0%/%1%/%3", "%3", "jono"],
				["Dashna Anand", "%cmo", "%0%/%1%/%4", "%4", "%4"]
			  ]
			]
		  }`,
		Subs: `{
			"ceo": "Chief Executive Officer",
			"cto": "Chief Technology Officer",
			"cmo": "Chief Marketing Officer"
		  }`,
		Trans: `{
			"o": {
			  "assignKeys": ["n", "s", "w", "e"],
			  "replacePair": {
				"name": "%n",
				"strapline": "%s",
				"website": "%w",
				"employees": "%e"
			  }
			},
			"e" : {
			  "arrayItems" : "em"
			},
			"em": {
			  "assignKeys": ["n", "p", "b", "l", "t"],
			  "replacePair": {
				"name": "%n",
				"position": "%p",
				"bio": "%b",
				"linkedin": "https://www.linkedin.com/in/%l",
				"twitter": "https://www.twitter.com/%t"
			  }
			}
		  }`,
		Output: `{
			"name": "Widget Company Ltd",
			"strapline": "Making the best widgets",
			"website": "https://www.widgetcompany.com",
			"employees": [{
			  "name": "Jane Smith",
			  "position": "Chief Executive Officer",
			  "bio": "https://www.widgetcompany.com/team/janesmith",
			  "linkedin": "https://www.linkedin.com/in/janesmith",
			  "twitter": "https://www.twitter.com/janesmithwidgets"
			}, {
			  "name": "John Wilson",
			  "position": "Chief Technology Officer",
			  "bio": "https://www.widgetcompany.com/team/johnwilson",
			  "linkedin": "https://www.linkedin.com/in/johnwilson",
			  "twitter": "https://www.twitter.com/jono"
			}, {
			  "name": "Dashna Anand",
			  "position": "Chief Marketing Officer",
			  "bio": "https://www.widgetcompany.com/team/dashnaanand",
			  "linkedin": "https://www.linkedin.com/in/dashnaanand",
			  "twitter": "https://www.twitter.com/dashnaanand"
			}]
		  }`,
	}, {
		Name:   "rewriteKey",
		Input:  `{"f": 1}`,
		Trans:  `{"f": {"rewriteKey": "foo"}}`,
		Output: `{"foo": 1}`,
	}, {
		Name:   "rewriteValue",
		Input:  `{"foo": 1}`,
		Trans:  `{"foo": {"rewriteValue" : { "x" : "%self", "y" : 2 }}}`,
		Output: `{"foo": {"x": 1, "y": 2}}`,
	}, {
		Name:   "replacePair",
		Input:  `{"f": 1}`,
		Trans:  `{"f": {"replacePair": {"x": "%self", "y": 2, "z": 3}}}`,
		Output: `{"x" : 1, "y" : 2, "z" : 3}`,
	}, {
		Name:   "subs",
		Input:  `{"foo": "%var"}`,
		Subs:   `{"var": 1}`,
		Output: `{"foo": 1}`,
	}, {
		Name:   "assignKeys",
		Input:  `{"phonetic": ["alpha", "bravo", "charlie"]}`,
		Trans:  `{"phonetic": {"assignKeys": ["a", "b", "c"]}}`,
		Output: `{"phonetic": {"a": "alpha", "b": "bravo", "c": "charlie"}}`,
	}, {
		Name:  "arrayItems",
		Input: `{"employees": [["Jane Smith", "CEO"], ["Dashna Anand", "CMO"]]}`,
		Trans: `{
			"employees": {"arrayItems": "employee"},
			"employee": {
				"assignKeys": ["name", "position"],
				"replacePair": {"name": "%name", "position": "%position"}
			}}`,
		Output: `{
			"employees": [
				{"name": "Jane Smith", "position": "CEO"},
				{"name": "Dashna Anand", "position": "CMO"}
			]}`,
	}, {
		Name: "noop",
		Input: `{
			"employees": [
				{"name": "Jane Smith", "position": "CEO"},
				{"name": "Dashna Anand", "position": "CMO"}
			]}`,
		Output: `{
			"employees": [
				{"name": "Jane Smith", "position": "CEO"},
				{"name": "Dashna Anand", "position": "CMO"}
			]}`,
	},
}

type UnpackerTest struct {
	Name   string
	Input  string
	Subs   string
	Trans  string
	Output string
	Skip   bool
}

func (test *UnpackerTest) Run(t *testing.T) {
	if testing.Short() && test.Skip {
		t.Skip("manually disabled test")
	}
	u := &Unpacker{}
	if test.Trans != `` {
		trans, err := ParseTransforms([]byte(test.Trans))
		if err != nil {
			t.Fatalf("Unable to parse transforms: %v", err)
		}
		u.Transforms = trans
	}
	if test.Subs != `` {
		subs := make(Substitution)
		if err := json.Unmarshal([]byte(test.Subs), &subs); err != nil {
			t.Fatalf("Unable to parse substition: %v", err)
		}
		u.AddSubs(subs)
	}
	output, err := u.Unpack([]byte(test.Input))
	if err != nil {
		t.Fatalf("Unable to unpack: %v", err)
	}
	exp := test.out()
	if !bytes.Equal(output, exp) {
		t.Fatalf("Unexpected output:\n\tWant: %s\n\tGot:  %s", string(exp), string(output))
	}
}

func (test UnpackerTest) out() []byte {
	var obj interface{}
	err := json.Unmarshal([]byte(test.Output), &obj)
	if err != nil {
		panic(err)
	}
	bits, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return bits
}

func TestUnpacker(t *testing.T) {
	for _, test := range unpackerTests {
		t.Run(test.Name, test.Run)
	}
}
