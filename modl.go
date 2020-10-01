// Package modl implements encoding and decoding of MODL as defined in modl.uk.
package modl

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/bign8/modl.go/internal/fs"
	"github.com/bign8/modl.go/internal/parser"
)

//go:generate java --version # assert java exists (sudo apt install default-jdk)
//go:generate wget --no-clobber https://www.antlr.org/download/antlr-4.8-complete.jar
//go:generate java -jar ./antlr-4.8-complete.jar -Dlanguage=Go -Xexact-output-dir -o internal/parser grammar/antlr4/MODLLexer.g4
//go:generate java -jar ./antlr-4.8-complete.jar -Dlanguage=Go -Xexact-output-dir -o internal/parser grammar/antlr4/MODLParser.g4

// Unmarshal parses the MODL-encoded data and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}, files fs.FS) error {
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
	Names []string
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

func (u *unmarshaler) EnterEveryRule(ctx antlr.ParserRuleContext) {
	println(u.Names[ctx.GetRuleIndex()] + "\t" + ctx.GetText())
}

func (u *unmarshaler) EnterModl_map(ctx *parser.Modl_mapContext) {
	println("Manually injecting map (FUTURE: respect provided fields)")
	m := map[string]interface{}{}
	value := reflect.ValueOf(m)
	u.push(value)
}

func (u *unmarshaler) ExitModl_array_item(ctx *parser.Modl_array_itemContext) {
	value := u.pop()
	ptr := u.peek() // TODO: assert is an pointer to an interface to an array
	if ptr.Kind() != reflect.Ptr {
		println("Exit Array (ptr): " + ptr.String())
		return // faking it
	}
	itr := ptr.Elem()
	if itr.Kind() != reflect.Interface {
		println("Exit Array (itr): " + itr.String())
		return // faking it
	}
	arr := itr.Elem()
	if arr.Kind() != reflect.Slice {
		println("Exit Array (arr): " + arr.Kind().String())
		return // faking it
	}
	arr = reflect.Append(arr, value)
	itr.Set(arr)
}

func (u *unmarshaler) ExitModl_pair(ctx *parser.Modl_pairContext) {
	value := u.pop()        // just finished parsing
	v := indirect(u.peek()) // get to a solid footing
	node := ctx.STRING()
	if node == nil {
		node = ctx.QUOTED()
	}
	key := node.GetText()
	if len(key) > 0 && key[0] == '*' {
		println("INSTRUCTION, ignoring (for now): " + key)
		return
	}
	key = decode(key)
	switch v.Kind() {
	case reflect.Map:
		v.SetMapIndex(reflect.ValueOf(key), value)
	default:
		println("What is this! " + v.Kind().String())
	}
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
		u.push(reflect.ValueOf(decode(ctx.STRING().GetText())))
	case parser.MODLParserQUOTED:
		text := ctx.QUOTED().GetText()
		if strings.HasPrefix(text, "`") && strings.HasSuffix(text, "`") {
			text = text[1 : len(text)-1]
		} else if strings.HasPrefix(text, "\"") && strings.HasSuffix(text, "\"") {
			text = text[1 : len(text)-1]
		}
		u.push(reflect.ValueOf(decode(text)))
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

func indirect(v reflect.Value) reflect.Value {
	// TODO: set nil types if not interfaces (see encoding/json/decode.go:indirect)
	for {
		switch v.Kind() {
		case reflect.Ptr, reflect.Interface:
			v = v.Elem()
		default:
			return v
		}
	}
}

func decode(in string) string {
	runes := []rune(in)
	j := 0
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if (r == '\\' || r == '~') && runes[i+1] == 'u' {
			if rx, off := getu4(runes[i:]); rx != -1 {
				runes[j] = rx
				i += off
				j++
				continue
			}
		}
		if r == '\\' || r == '~' {
			r = runes[i+1]
			i++
		}
		runes[j] = r
		j++
	}
	return string(runes[:j])
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
