// Code generated from grammar/antlr4/MODLParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MODLParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMODLParserListener is a complete listener for a parse tree produced by MODLParser.
type BaseMODLParserListener struct{}

var _ MODLParserListener = &BaseMODLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMODLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMODLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMODLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMODLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModl is called when production modl is entered.
func (s *BaseMODLParserListener) EnterModl(ctx *ModlContext) {}

// ExitModl is called when production modl is exited.
func (s *BaseMODLParserListener) ExitModl(ctx *ModlContext) {}

// EnterModl_structure is called when production modl_structure is entered.
func (s *BaseMODLParserListener) EnterModl_structure(ctx *Modl_structureContext) {}

// ExitModl_structure is called when production modl_structure is exited.
func (s *BaseMODLParserListener) ExitModl_structure(ctx *Modl_structureContext) {}

// EnterModl_map is called when production modl_map is entered.
func (s *BaseMODLParserListener) EnterModl_map(ctx *Modl_mapContext) {}

// ExitModl_map is called when production modl_map is exited.
func (s *BaseMODLParserListener) ExitModl_map(ctx *Modl_mapContext) {}

// EnterModl_array is called when production modl_array is entered.
func (s *BaseMODLParserListener) EnterModl_array(ctx *Modl_arrayContext) {}

// ExitModl_array is called when production modl_array is exited.
func (s *BaseMODLParserListener) ExitModl_array(ctx *Modl_arrayContext) {}

// EnterModl_pair is called when production modl_pair is entered.
func (s *BaseMODLParserListener) EnterModl_pair(ctx *Modl_pairContext) {}

// ExitModl_pair is called when production modl_pair is exited.
func (s *BaseMODLParserListener) ExitModl_pair(ctx *Modl_pairContext) {}

// EnterModl_value is called when production modl_value is entered.
func (s *BaseMODLParserListener) EnterModl_value(ctx *Modl_valueContext) {}

// ExitModl_value is called when production modl_value is exited.
func (s *BaseMODLParserListener) ExitModl_value(ctx *Modl_valueContext) {}

// EnterModl_primitive is called when production modl_primitive is entered.
func (s *BaseMODLParserListener) EnterModl_primitive(ctx *Modl_primitiveContext) {}

// ExitModl_primitive is called when production modl_primitive is exited.
func (s *BaseMODLParserListener) ExitModl_primitive(ctx *Modl_primitiveContext) {}
