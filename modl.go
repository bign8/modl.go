// Package modl implements encoding and decoding of MODL as defined in modl.uk.
package modl

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/bign8/modl.go/internal/fs"
	"github.com/bign8/modl.go/internal/parser"
)

//go:generate java --version # assert java exists (sudo apt install default-jdk)
//go:generate wget --no-clobber https://www.antlr.org/download/antlr-4.8-complete.jar
//go:generate java -jar ./antlr-4.8-complete.jar -Dlanguage=Go -Xexact-output-dir -o internal/parser grammar/antlr4/MODLLexer.g4
//go:generate rm internal/parser/MODLLexer.interp internal/parser/MODLLexer.tokens

// Unmarshal parses the MODL-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(data []byte, v interface{}, files fs.FS) error {
	is := antlr.NewInputStream(string(data))
	lexer := parser.NewMODLLexer(is)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
	return nil // TODO: everything :cry:
}
