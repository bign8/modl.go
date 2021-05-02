// Code generated from grammar/antlr4/MODLParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MODLParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 81, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 7, 2, 20, 10, 2, 12, 2, 14, 2, 23, 11, 2, 3,
	2, 5, 2, 26, 10, 2, 3, 2, 5, 2, 29, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	5, 3, 36, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 42, 10, 4, 12, 4, 14, 4,
	45, 11, 4, 5, 4, 47, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 55,
	10, 5, 12, 5, 14, 5, 58, 11, 5, 5, 5, 60, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 71, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	5, 7, 77, 10, 7, 3, 8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2,
	4, 3, 2, 15, 16, 4, 2, 4, 6, 14, 16, 2, 87, 2, 28, 3, 2, 2, 2, 4, 35, 3,
	2, 2, 2, 6, 37, 3, 2, 2, 2, 8, 50, 3, 2, 2, 2, 10, 70, 3, 2, 2, 2, 12,
	76, 3, 2, 2, 2, 14, 78, 3, 2, 2, 2, 16, 21, 5, 4, 3, 2, 17, 18, 7, 8, 2,
	2, 18, 20, 5, 4, 3, 2, 19, 17, 3, 2, 2, 2, 20, 23, 3, 2, 2, 2, 21, 19,
	3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2,
	24, 26, 7, 8, 2, 2, 25, 24, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 29, 3,
	2, 2, 2, 27, 29, 5, 14, 8, 2, 28, 16, 3, 2, 2, 2, 28, 27, 3, 2, 2, 2, 29,
	30, 3, 2, 2, 2, 30, 31, 7, 2, 2, 3, 31, 3, 3, 2, 2, 2, 32, 36, 5, 6, 4,
	2, 33, 36, 5, 8, 5, 2, 34, 36, 5, 10, 6, 2, 35, 32, 3, 2, 2, 2, 35, 33,
	3, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 5, 3, 2, 2, 2, 37, 46, 7, 10, 2, 2,
	38, 43, 5, 10, 6, 2, 39, 40, 7, 8, 2, 2, 40, 42, 5, 10, 6, 2, 41, 39, 3,
	2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44,
	47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 38, 3, 2, 2, 2, 46, 47, 3, 2, 2,
	2, 47, 48, 3, 2, 2, 2, 48, 49, 7, 11, 2, 2, 49, 7, 3, 2, 2, 2, 50, 59,
	7, 12, 2, 2, 51, 56, 5, 12, 7, 2, 52, 53, 7, 8, 2, 2, 53, 55, 5, 12, 7,
	2, 54, 52, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57,
	3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 51, 3, 2, 2, 2,
	59, 60, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 62, 7, 13, 2, 2, 62, 9, 3,
	2, 2, 2, 63, 64, 9, 2, 2, 2, 64, 65, 7, 7, 2, 2, 65, 71, 5, 12, 7, 2, 66,
	67, 7, 16, 2, 2, 67, 71, 5, 6, 4, 2, 68, 69, 7, 16, 2, 2, 69, 71, 5, 8,
	5, 2, 70, 63, 3, 2, 2, 2, 70, 66, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 11,
	3, 2, 2, 2, 72, 77, 5, 6, 4, 2, 73, 77, 5, 8, 5, 2, 74, 77, 5, 10, 6, 2,
	75, 77, 5, 14, 8, 2, 76, 72, 3, 2, 2, 2, 76, 73, 3, 2, 2, 2, 76, 74, 3,
	2, 2, 2, 76, 75, 3, 2, 2, 2, 77, 13, 3, 2, 2, 2, 78, 79, 9, 3, 2, 2, 79,
	15, 3, 2, 2, 2, 12, 21, 25, 28, 35, 43, 46, 56, 59, 70, 76,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'null'", "'true'", "'false'", "'='", "';'", "','", "'('", "')'",
	"'['", "']'",
}
var symbolicNames = []string{
	"", "WS", "NULL", "TRUE", "FALSE", "EQUALS", "STRUCT_SEP", "ARR_SEP", "LBRAC",
	"RBRAC", "LSBRAC", "RSBRAC", "NUMBER", "QUOTED", "STRING",
}

var ruleNames = []string{
	"modl", "modl_structure", "modl_map", "modl_array", "modl_pair", "modl_value",
	"modl_primitive",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MODLParser struct {
	*antlr.BaseParser
}

func NewMODLParser(input antlr.TokenStream) *MODLParser {
	this := new(MODLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "MODLParser.g4"

	return this
}

// MODLParser tokens.
const (
	MODLParserEOF        = antlr.TokenEOF
	MODLParserWS         = 1
	MODLParserNULL       = 2
	MODLParserTRUE       = 3
	MODLParserFALSE      = 4
	MODLParserEQUALS     = 5
	MODLParserSTRUCT_SEP = 6
	MODLParserARR_SEP    = 7
	MODLParserLBRAC      = 8
	MODLParserRBRAC      = 9
	MODLParserLSBRAC     = 10
	MODLParserRSBRAC     = 11
	MODLParserNUMBER     = 12
	MODLParserQUOTED     = 13
	MODLParserSTRING     = 14
)

// MODLParser rules.
const (
	MODLParserRULE_modl           = 0
	MODLParserRULE_modl_structure = 1
	MODLParserRULE_modl_map       = 2
	MODLParserRULE_modl_array     = 3
	MODLParserRULE_modl_pair      = 4
	MODLParserRULE_modl_value     = 5
	MODLParserRULE_modl_primitive = 6
)

// IModlContext is an interface to support dynamic dispatch.
type IModlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModlContext differentiates from other interfaces.
	IsModlContext()
}

type ModlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModlContext() *ModlContext {
	var p = new(ModlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl
	return p
}

func (*ModlContext) IsModlContext() {}

func NewModlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModlContext {
	var p = new(ModlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl

	return p
}

func (s *ModlContext) GetParser() antlr.Parser { return s.parser }

func (s *ModlContext) EOF() antlr.TerminalNode {
	return s.GetToken(MODLParserEOF, 0)
}

func (s *ModlContext) Modl_primitive() IModl_primitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_primitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_primitiveContext)
}

func (s *ModlContext) AllModl_structure() []IModl_structureContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_structureContext)(nil)).Elem())
	var tst = make([]IModl_structureContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_structureContext)
		}
	}

	return tst
}

func (s *ModlContext) Modl_structure(i int) IModl_structureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_structureContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_structureContext)
}

func (s *ModlContext) AllSTRUCT_SEP() []antlr.TerminalNode {
	return s.GetTokens(MODLParserSTRUCT_SEP)
}

func (s *ModlContext) STRUCT_SEP(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserSTRUCT_SEP, i)
}

func (s *ModlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl(s)
	}
}

func (s *ModlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl(s)
	}
}

func (p *MODLParser) Modl() (localctx IModlContext) {
	localctx = NewModlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MODLParserRULE_modl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(14)
			p.Modl_structure()
		}
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(15)
					p.Match(MODLParserSTRUCT_SEP)
				}
				{
					p.SetState(16)
					p.Modl_structure()
				}

			}
			p.SetState(21)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MODLParserSTRUCT_SEP {
			{
				p.SetState(22)
				p.Match(MODLParserSTRUCT_SEP)
			}

		}

	case 2:
		{
			p.SetState(25)
			p.Modl_primitive()
		}

	}
	{
		p.SetState(28)
		p.Match(MODLParserEOF)
	}

	return localctx
}

// IModl_structureContext is an interface to support dynamic dispatch.
type IModl_structureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_structureContext differentiates from other interfaces.
	IsModl_structureContext()
}

type Modl_structureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_structureContext() *Modl_structureContext {
	var p = new(Modl_structureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_structure
	return p
}

func (*Modl_structureContext) IsModl_structureContext() {}

func NewModl_structureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_structureContext {
	var p = new(Modl_structureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_structure

	return p
}

func (s *Modl_structureContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_structureContext) Modl_map() IModl_mapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_mapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_mapContext)
}

func (s *Modl_structureContext) Modl_array() IModl_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_arrayContext)
}

func (s *Modl_structureContext) Modl_pair() IModl_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_pairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_pairContext)
}

func (s *Modl_structureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_structureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_structureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_structure(s)
	}
}

func (s *Modl_structureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_structure(s)
	}
}

func (p *MODLParser) Modl_structure() (localctx IModl_structureContext) {
	localctx = NewModl_structureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MODLParserRULE_modl_structure)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MODLParserLBRAC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Modl_map()
		}

	case MODLParserLSBRAC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Modl_array()
		}

	case MODLParserQUOTED, MODLParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)
			p.Modl_pair()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IModl_mapContext is an interface to support dynamic dispatch.
type IModl_mapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_mapContext differentiates from other interfaces.
	IsModl_mapContext()
}

type Modl_mapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_mapContext() *Modl_mapContext {
	var p = new(Modl_mapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_map
	return p
}

func (*Modl_mapContext) IsModl_mapContext() {}

func NewModl_mapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_mapContext {
	var p = new(Modl_mapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_map

	return p
}

func (s *Modl_mapContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_mapContext) LBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserLBRAC, 0)
}

func (s *Modl_mapContext) RBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserRBRAC, 0)
}

func (s *Modl_mapContext) AllModl_pair() []IModl_pairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_pairContext)(nil)).Elem())
	var tst = make([]IModl_pairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_pairContext)
		}
	}

	return tst
}

func (s *Modl_mapContext) Modl_pair(i int) IModl_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_pairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_pairContext)
}

func (s *Modl_mapContext) AllSTRUCT_SEP() []antlr.TerminalNode {
	return s.GetTokens(MODLParserSTRUCT_SEP)
}

func (s *Modl_mapContext) STRUCT_SEP(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserSTRUCT_SEP, i)
}

func (s *Modl_mapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_mapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_mapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_map(s)
	}
}

func (s *Modl_mapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_map(s)
	}
}

func (p *MODLParser) Modl_map() (localctx IModl_mapContext) {
	localctx = NewModl_mapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MODLParserRULE_modl_map)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(MODLParserLBRAC)
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MODLParserQUOTED || _la == MODLParserSTRING {
		{
			p.SetState(36)
			p.Modl_pair()
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MODLParserSTRUCT_SEP {
			{
				p.SetState(37)
				p.Match(MODLParserSTRUCT_SEP)
			}
			{
				p.SetState(38)
				p.Modl_pair()
			}

			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(46)
		p.Match(MODLParserRBRAC)
	}

	return localctx
}

// IModl_arrayContext is an interface to support dynamic dispatch.
type IModl_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_arrayContext differentiates from other interfaces.
	IsModl_arrayContext()
}

type Modl_arrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_arrayContext() *Modl_arrayContext {
	var p = new(Modl_arrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_array
	return p
}

func (*Modl_arrayContext) IsModl_arrayContext() {}

func NewModl_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_arrayContext {
	var p = new(Modl_arrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_array

	return p
}

func (s *Modl_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_arrayContext) LSBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserLSBRAC, 0)
}

func (s *Modl_arrayContext) RSBRAC() antlr.TerminalNode {
	return s.GetToken(MODLParserRSBRAC, 0)
}

func (s *Modl_arrayContext) AllModl_value() []IModl_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModl_valueContext)(nil)).Elem())
	var tst = make([]IModl_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModl_valueContext)
		}
	}

	return tst
}

func (s *Modl_arrayContext) Modl_value(i int) IModl_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModl_valueContext)
}

func (s *Modl_arrayContext) AllSTRUCT_SEP() []antlr.TerminalNode {
	return s.GetTokens(MODLParserSTRUCT_SEP)
}

func (s *Modl_arrayContext) STRUCT_SEP(i int) antlr.TerminalNode {
	return s.GetToken(MODLParserSTRUCT_SEP, i)
}

func (s *Modl_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_array(s)
	}
}

func (s *Modl_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_array(s)
	}
}

func (p *MODLParser) Modl_array() (localctx IModl_arrayContext) {
	localctx = NewModl_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MODLParserRULE_modl_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(MODLParserLSBRAC)
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserNULL)|(1<<MODLParserTRUE)|(1<<MODLParserFALSE)|(1<<MODLParserLBRAC)|(1<<MODLParserLSBRAC)|(1<<MODLParserNUMBER)|(1<<MODLParserQUOTED)|(1<<MODLParserSTRING))) != 0 {
		{
			p.SetState(49)
			p.Modl_value()
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MODLParserSTRUCT_SEP {
			{
				p.SetState(50)
				p.Match(MODLParserSTRUCT_SEP)
			}
			{
				p.SetState(51)
				p.Modl_value()
			}

			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(59)
		p.Match(MODLParserRSBRAC)
	}

	return localctx
}

// IModl_pairContext is an interface to support dynamic dispatch.
type IModl_pairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_pairContext differentiates from other interfaces.
	IsModl_pairContext()
}

type Modl_pairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_pairContext() *Modl_pairContext {
	var p = new(Modl_pairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_pair
	return p
}

func (*Modl_pairContext) IsModl_pairContext() {}

func NewModl_pairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_pairContext {
	var p = new(Modl_pairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_pair

	return p
}

func (s *Modl_pairContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_pairContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(MODLParserEQUALS, 0)
}

func (s *Modl_pairContext) Modl_value() IModl_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_valueContext)
}

func (s *Modl_pairContext) STRING() antlr.TerminalNode {
	return s.GetToken(MODLParserSTRING, 0)
}

func (s *Modl_pairContext) QUOTED() antlr.TerminalNode {
	return s.GetToken(MODLParserQUOTED, 0)
}

func (s *Modl_pairContext) Modl_map() IModl_mapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_mapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_mapContext)
}

func (s *Modl_pairContext) Modl_array() IModl_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_arrayContext)
}

func (s *Modl_pairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_pairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_pairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_pair(s)
	}
}

func (s *Modl_pairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_pair(s)
	}
}

func (p *MODLParser) Modl_pair() (localctx IModl_pairContext) {
	localctx = NewModl_pairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MODLParserRULE_modl_pair)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MODLParserQUOTED || _la == MODLParserSTRING) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(62)
			p.Match(MODLParserEQUALS)
		}
		{
			p.SetState(63)
			p.Modl_value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Match(MODLParserSTRING)
		}
		{
			p.SetState(65)
			p.Modl_map()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Match(MODLParserSTRING)
		}
		{
			p.SetState(67)
			p.Modl_array()
		}

	}

	return localctx
}

// IModl_valueContext is an interface to support dynamic dispatch.
type IModl_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_valueContext differentiates from other interfaces.
	IsModl_valueContext()
}

type Modl_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_valueContext() *Modl_valueContext {
	var p = new(Modl_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_value
	return p
}

func (*Modl_valueContext) IsModl_valueContext() {}

func NewModl_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_valueContext {
	var p = new(Modl_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_value

	return p
}

func (s *Modl_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_valueContext) Modl_map() IModl_mapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_mapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_mapContext)
}

func (s *Modl_valueContext) Modl_array() IModl_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_arrayContext)
}

func (s *Modl_valueContext) Modl_pair() IModl_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_pairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_pairContext)
}

func (s *Modl_valueContext) Modl_primitive() IModl_primitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModl_primitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModl_primitiveContext)
}

func (s *Modl_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_value(s)
	}
}

func (s *Modl_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_value(s)
	}
}

func (p *MODLParser) Modl_value() (localctx IModl_valueContext) {
	localctx = NewModl_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MODLParserRULE_modl_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Modl_map()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Modl_array()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.Modl_pair()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(73)
			p.Modl_primitive()
		}

	}

	return localctx
}

// IModl_primitiveContext is an interface to support dynamic dispatch.
type IModl_primitiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModl_primitiveContext differentiates from other interfaces.
	IsModl_primitiveContext()
}

type Modl_primitiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModl_primitiveContext() *Modl_primitiveContext {
	var p = new(Modl_primitiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MODLParserRULE_modl_primitive
	return p
}

func (*Modl_primitiveContext) IsModl_primitiveContext() {}

func NewModl_primitiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modl_primitiveContext {
	var p = new(Modl_primitiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MODLParserRULE_modl_primitive

	return p
}

func (s *Modl_primitiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Modl_primitiveContext) QUOTED() antlr.TerminalNode {
	return s.GetToken(MODLParserQUOTED, 0)
}

func (s *Modl_primitiveContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MODLParserNUMBER, 0)
}

func (s *Modl_primitiveContext) STRING() antlr.TerminalNode {
	return s.GetToken(MODLParserSTRING, 0)
}

func (s *Modl_primitiveContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MODLParserTRUE, 0)
}

func (s *Modl_primitiveContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MODLParserFALSE, 0)
}

func (s *Modl_primitiveContext) NULL() antlr.TerminalNode {
	return s.GetToken(MODLParserNULL, 0)
}

func (s *Modl_primitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modl_primitiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modl_primitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.EnterModl_primitive(s)
	}
}

func (s *Modl_primitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MODLParserListener); ok {
		listenerT.ExitModl_primitive(s)
	}
}

func (p *MODLParser) Modl_primitive() (localctx IModl_primitiveContext) {
	localctx = NewModl_primitiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MODLParserRULE_modl_primitive)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MODLParserNULL)|(1<<MODLParserTRUE)|(1<<MODLParserFALSE)|(1<<MODLParserNUMBER)|(1<<MODLParserQUOTED)|(1<<MODLParserSTRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
