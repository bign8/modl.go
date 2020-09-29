// Package modl implements encoding and decoding of MODL as defined in modl.uk.
package modl

import (
	"errors"
	"reflect"

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

func (u *unmarshaler) ExitModl_pair(ctx *parser.Modl_pairContext) {
	value := u.pop()        // just finished parsing
	v := indirect(u.peek()) // get to a solid footing
	switch v.Kind() {
	case reflect.Map:
		if ctx.STRING() == nil {
			return // TODO: QUOTED case
		}
		key := ctx.STRING().GetText()
		v.SetMapIndex(reflect.ValueOf(key), value)
	default:
		println("What is this! " + v.Kind().String())
	}
	println("Exiting pair " + ctx.STRING().GetText() + " " + indirect(v).Kind().String())
}

func (u *unmarshaler) ExitModl_value(ctx *parser.Modl_valueContext) {
	value := ctx.GetText() // not correct!!! can be a map, array, nb_array, pair or primitive!
	u.push(reflect.ValueOf(value))
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
