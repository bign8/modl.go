// Package modl implements encoding and decoding of MODL as defined in modl.uk.
package modl

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/bign8/modl.go/internal/parser"
)

//go:generate java --version # assert java exists (sudo apt install default-jdk)
//go:generate wget --no-clobber https://www.antlr.org/download/antlr-4.8-complete.jar
//go:generate java -jar ./antlr-4.8-complete.jar -Dlanguage=Go -Xexact-output-dir -o internal/parser grammar/antlr4/MODLLexer.g4
//go:generate java -jar ./antlr-4.8-complete.jar -Dlanguage=Go -Xexact-output-dir -o internal/parser grammar/antlr4/MODLParser.g4

// Unmarshal parses the MODL-encoded data and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("modl: Unmarshal(" + reflect.TypeOf(v).Kind().String() + ")")
	}
	is := antlr.NewInputStream(string(data))
	lexer := parser.NewMODLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewMODLParser(stream)
	state := &unmarshaler{Names: p.RuleNames}
	state.push(rv)
	antlr.ParseTreeWalkerDefault.Walk(state, p.Modl())
	return state.err
}

type unmarshaler struct {
	*parser.BaseMODLParserListener
	Names []string // DEBUG ONLY (REMOVE ON RELEASE PLZ!!!)
	stack []reflect.Value
	err   error
}

func (u *unmarshaler) peek() reflect.Value {
	if len(u.stack) < 1 {
		return reflect.Value{ /* "Soft" panic :cry: */ }
	}
	return u.stack[len(u.stack)-1]
}

func (u *unmarshaler) push(v reflect.Value) { u.stack = append(u.stack, v) }
func (u *unmarshaler) pop() reflect.Value {
	if len(u.stack) < 1 {
		return reflect.Value{ /* "Soft" panic :cry: */ }
	}
	tail := len(u.stack) - 1
	result := u.stack[tail]
	u.stack = u.stack[:tail]
	return result
}
func (u *unmarshaler) debug() {
	fmt.Println("Stack Start ............ ") // use fmt so I don't have to keep swapping imports :cry:
	for i, item := range u.stack {
		println("Stack Item(" + strconv.Itoa(i) + "): " + item.String())
	}
	println("DEBUG END ============== ")
}

// func (u *unmarshaler) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	println(u.Names[ctx.GetRuleIndex()] + "\tEnter\t" + ctx.GetText())
// }

// func (u *unmarshaler) ExitEveryRule(ctx antlr.ParserRuleContext) {
// 	println(u.Names[ctx.GetRuleIndex()] + "\tExit\t" + ctx.GetText())
// }

func (u *unmarshaler) EnterModl(ctx *parser.ModlContext) {
	// Recurse through some externally provided interfaces (how I do tests :cry:)
	//   x := map[string]interface{}{}
	//   var y interface{}
	//   y = x
	//   Unmarshal(data, &y, nil)
}

func (u *unmarshaler) ExitModl(ctx *parser.ModlContext) {
	// Verify the stack is still valid (should generally be a noop)
	for _, v := range u.stack {
		if !v.IsValid() {
			u.debug()
			println("TODO(fixme): invalid in Stack")
			return
		}
	}

	// Iff we have more than 2 items (pointer and structure), we must have "orphan pairs"
	// pairs outside of a map at the top level are considered pairs in the same map
	if len(u.stack) > 2 {
		// Merge subsequent orphan pairs, onto first orphan
		base := u.stack[1]
		for _, orphan := range u.stack[2:] {
			if orphan.Kind() != reflect.Map {
				panic("parent orphan that wasn't a map!")
			}
			key := orphan.MapKeys()[0] // pairs should only have one key
			value := orphan.MapIndex(key)
			base.SetMapIndex(key, value)
		}
		u.stack = u.stack[:2] // slice stack back down to size
	}

	// Otherwise, we have 2 items, pop the "value" and assign to "ptr"
	v := u.pop()
	u.peek().Elem().Set(v)
}

func (u *unmarshaler) EnterModl_map(ctx *parser.Modl_mapContext) {

	// Fix for b216, the breaks b049, b050, b064 without the depth qualifier
	if len(ctx.AllModl_pair()) == 0 && len(u.stack) == 1 {
		var m map[string]interface{}
		value := reflect.ValueOf(m)
		u.push(value)
		return
	}

	// println("Manually injecting map (TODO: respect provided fields)")
	m := map[string]interface{}{}
	value := reflect.ValueOf(m)
	u.push(value)
}

func (u *unmarshaler) EnterModl_array(ctx *parser.Modl_arrayContext) {
	a := make([]interface{}, 0, len(ctx.AllModl_value()))
	u.push(reflect.ValueOf(a))
}

func (u *unmarshaler) ExitModl_array(ctx *parser.Modl_arrayContext) {
	seps := ctx.AllSTRUCT_SEP()
	markers := make([]int, len(seps))
	for i, sep := range seps {
		markers[i] = sep.GetSymbol().GetStart()
	}
	ptr := len(u.stack) - len(ctx.AllModl_value())
	if ptr < 1 {
		panic("exitArray: invalid stack... gtfo") // TODO: exit cleanly
	}
	items := u.stack[ptr:]
	arr := u.stack[ptr-1]
	if arr.Kind() != reflect.Slice {
		println("exitArray: not a slice... skipping: " + arr.Kind().String())
		return
	}

	// Inject NULLs for adjacent markers (array item separators)
	for i := 1; i < len(markers); i++ {
		if markers[i-1]+1 == markers[i] {
			items = append(items, reflect.Value{})    // make space for null value
			copy(items[i+1:], items[i:])              // shifting elements
			items[i] = reflect.New(arr.Type().Elem()) // Null item (based on array type)
		}
	}
	u.stack[ptr-1] = reflect.Append(arr, items...) // Append the elements into the slice
	u.stack = u.stack[:ptr]                        // slice off the items from the stack
}

func (u *unmarshaler) EnterModl_pair(ctx *parser.Modl_pairContext) {

	// See if parent has a property that can be set, matching our key
	v0 := u.peek()
	if v0.IsValid() {
		v := indirect(v0, false) // ensure we are on solid ground
		switch v.Kind() {
		case reflect.Map:
			// we "likely" support setting this as a key, don't overwrite existing map
			return
		}
	}

	// put an assignable PAIR onto the stack
	m := map[string]interface{}{}
	value := reflect.ValueOf(m)
	u.push(value)
}

func (u *unmarshaler) ExitModl_pair(ctx *parser.Modl_pairContext) {
	value := u.pop() // just finished parsing
	if !u.peek().IsValid() {
		println("ExitModl_pair: Unexpected state... gtfo")
		return
	}
	v := indirect(u.peek(), false) // get to a solid footing
	node := ctx.STRING()
	if node == nil {
		node = ctx.QUOTED()
	}
	key := node.GetText()
	key = decode(key).String()

	// Just set the key=value
	switch v.Kind() {
	case reflect.Map:
		v.SetMapIndex(reflect.ValueOf(key), value)
	default:
		u.push(value) // put value back
		println("ExitModl_pair: What is this: " + v.Kind().String())
	}
}

func (u *unmarshaler) EnterModl_value(ctx *parser.Modl_valueContext) {
	// Add an invalid type, in case we enter another pair (don't re-use parent objects)
	u.push(reflect.Value{})
}

func (u *unmarshaler) ExitModl_value(ctx *parser.Modl_valueContext) {
	// the thing we just made as a value
	v := u.pop()
	if !v.IsValid() {
		println("ExitModl_value(1): created value should be valid!")
	}

	// pop our invalid marker from the stack
	t := u.pop()
	if t.IsValid() {
		println("ExitModl_value(2): marker should be invalid: " + t.String())
		return
	}

	u.push(v)
}

func (u *unmarshaler) EnterModl_primitive(ctx *parser.Modl_primitiveContext) {
	switch ctx.GetStart().GetTokenType() {
	case parser.MODLParserNUMBER:
		f, err := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
		if err != nil {
			panic(err)
		}
		u.push(reflect.ValueOf(f))
	case parser.MODLParserSTRING:
		text := ctx.STRING().GetText()
		u.push(decode(text))
	case parser.MODLParserQUOTED:
		text := ctx.QUOTED().GetText()
		u.push(decode(text))
	case parser.MODLParserTRUE:
		u.push(reflect.ValueOf(true))
	case parser.MODLParserFALSE:
		u.push(reflect.ValueOf(false))
	case parser.MODLParserNULL:
		u.push(reflect.ValueOf((*interface{})(nil)))
	default:
		panic("UNKNOWN PRIMITIVE TYPE: " + ctx.GetText())
	}
}

// TODO: set nil types if not interfaces (see encoding/json/decode.go:indirect)
// indirect walks down v allocating pointers as needed,
// until it gets to a non-pointer.
// If it encounters an Unmarshaler, indirect stops and returns that.
// If decodingNull is true, indirect stops at the first settable pointer so it
// can be set to nil.
func indirect(v reflect.Value, decodingNull bool) reflect.Value {
	// Issue #24153 indicates that it is generally not a guaranteed property
	// that you may round-trip a reflect.Value by calling Value.Addr().Elem()
	// and expect the value to still be settable for values derived from
	// unexported embedded struct fields.
	//
	// The logic below effectively does this when it first addresses the value
	// (to satisfy possible pointer methods) and continues to dereference
	// subsequent pointers as necessary.
	//
	// After the first round-trip, we set v back to the original value to
	// preserve the original RW flags contained in reflect.Value.
	v0 := v
	haveAddr := false

	// If v is a named type and is addressable,
	// start with its address, so that if the type has pointer methods,
	// we find them.
	if v.Kind() != reflect.Ptr && v.Type().Name() != "" && v.CanAddr() {
		haveAddr = true
		v = v.Addr()
	}
	for {
		// Load value from interface, but only if the result will be
		// usefully addressable.
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Ptr && !e.IsNil() && (!decodingNull || e.Elem().Kind() == reflect.Ptr) {
				haveAddr = false
				v = e
				continue
			}
		}

		if v.Kind() != reflect.Ptr {
			break
		}

		if decodingNull && v.CanSet() {
			break
		}

		// Prevent infinite loop if v is an interface pointing to its own address:
		//     var v interface{}
		//     v = &v
		if v.Elem().Kind() == reflect.Interface && v.Elem().Elem() == v {
			v = v.Elem()
			break
		}
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		// TODO: Unmarshaling type casting

		if haveAddr {
			v = v0 // restore original value after round-trip Value.Addr().Elem()
			haveAddr = false
		} else {
			v = v.Elem()
		}
	}
	return v
}

func decode(in string) reflect.Value {
	if strings.HasPrefix(in, "`") && strings.HasSuffix(in, "`") {
		in = in[1 : len(in)-1]
	} else if strings.HasPrefix(in, "\"") && strings.HasSuffix(in, "\"") {
		in = in[1 : len(in)-1]
	}

	// noop short circuit
	if len(in) == 0 { // noop
		return reflect.ValueOf(in)
	}

	// Decode string (inspired by unquoteBytes in encoding/json/decode.go)
	runes := []rune(in)
	j := 0
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if (r == '\\' || r == '~') && i < len(runes)-1 {
			switch runes[i+1] {
			case 'u':
				if rx, off := getu4(runes[i:]); rx != -1 {
					runes[j] = rx
					i += off
					j++
					continue
				}
			case '\\', '~':
				r = runes[i+1]
				i++ // skip forward
			case 'n':
				r = '\n'
				i++ // skip forward
			case 't':
				r = '\t'
				i++ // skip forward
				// Missing: b, f, r, u from encoding/json (future bug report?)
			}
		}
		runes[j] = r
		j++
	}
	return reflect.ValueOf(string(runes[:j]))
}

// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
// or it returns -1.  Borrowed and modified from encoding/json/decode.go
func getu4(s []rune) (rune, int) {
	if len(s) < 6 || s[1] != 'u' || !(s[0] == '\\' || s[0] == '~') {
		return -1, -1
	}
	var r, q rune
	for i, c := range s[2:] {
		if i >= 5 {
			return r, 6 // longest unicode utf-8 is currently 6 bytes
		}
		switch {
		case '0' <= c && c <= '9':
			c = c - '0'
		case 'a' <= c && c <= 'f':
			c = c - 'a' + 10
		case 'A' <= c && c <= 'F':
			c = c - 'A' + 10
		default:
			if i > 3 { // if we have at least 4 digits
				return r, i + 1
			}
			return -1, -1
		}
		q = r*16 + rune(c)
		if !utf8.ValidRune(q) {
			return r, i - 1
		}
		r = q
	}
	return r, len(s)
}
