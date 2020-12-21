// Code generated from Cacl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Cacl

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 34, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 22, 10, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8,
	10, 2, 4, 3, 2, 5, 9, 3, 2, 3, 4, 2, 31, 2, 12, 3, 2, 2, 2, 4, 21, 3, 2,
	2, 2, 6, 23, 3, 2, 2, 2, 8, 27, 3, 2, 2, 2, 10, 31, 3, 2, 2, 2, 12, 13,
	9, 2, 2, 2, 13, 3, 3, 2, 2, 2, 14, 22, 3, 2, 2, 2, 15, 22, 7, 4, 2, 2,
	16, 17, 5, 8, 5, 2, 17, 18, 7, 10, 2, 2, 18, 19, 5, 4, 3, 2, 19, 22, 3,
	2, 2, 2, 20, 22, 5, 6, 4, 2, 21, 14, 3, 2, 2, 2, 21, 15, 3, 2, 2, 2, 21,
	16, 3, 2, 2, 2, 21, 20, 3, 2, 2, 2, 22, 5, 3, 2, 2, 2, 23, 24, 9, 3, 2,
	2, 24, 25, 5, 2, 2, 2, 25, 26, 9, 3, 2, 2, 26, 7, 3, 2, 2, 2, 27, 28, 7,
	3, 2, 2, 28, 29, 7, 7, 2, 2, 29, 30, 7, 4, 2, 2, 30, 9, 3, 2, 2, 2, 31,
	32, 5, 4, 3, 2, 32, 11, 3, 2, 2, 2, 3, 21,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "'+'", "'-'", "'='", "'*'", "'/'", "'\n'", "'\r?\n'",
}
var symbolicNames = []string{
	"", "Id", "Int", "Add", "Sub", "EQ", "CHENG", "CHU", "NL", "HUANHANG",
	"WS",
}

var ruleNames = []string{
	"op", "expr", "addVal", "setVal", "init",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CaclParser struct {
	*antlr.BaseParser
}

func NewCaclParser(input antlr.TokenStream) *CaclParser {
	this := new(CaclParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Cacl.g4"

	return this
}

// CaclParser tokens.
const (
	CaclParserEOF      = antlr.TokenEOF
	CaclParserId       = 1
	CaclParserInt      = 2
	CaclParserAdd      = 3
	CaclParserSub      = 4
	CaclParserEQ       = 5
	CaclParserCHENG    = 6
	CaclParserCHU      = 7
	CaclParserNL       = 8
	CaclParserHUANHANG = 9
	CaclParserWS       = 10
)

// CaclParser rules.
const (
	CaclParserRULE_op     = 0
	CaclParserRULE_expr   = 1
	CaclParserRULE_addVal = 2
	CaclParserRULE_setVal = 3
	CaclParserRULE_init   = 4
)

// IOpContext is an interface to support dynamic dispatch.
type IOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpContext differentiates from other interfaces.
	IsOpContext()
}

type OpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpContext() *OpContext {
	var p = new(OpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CaclParserRULE_op
	return p
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaclParserRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }

func (s *OpContext) Add() antlr.TerminalNode {
	return s.GetToken(CaclParserAdd, 0)
}

func (s *OpContext) Sub() antlr.TerminalNode {
	return s.GetToken(CaclParserSub, 0)
}

func (s *OpContext) EQ() antlr.TerminalNode {
	return s.GetToken(CaclParserEQ, 0)
}

func (s *OpContext) CHENG() antlr.TerminalNode {
	return s.GetToken(CaclParserCHENG, 0)
}

func (s *OpContext) CHU() antlr.TerminalNode {
	return s.GetToken(CaclParserCHU, 0)
}

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaclListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaclListener); ok {
		listenerT.ExitOp(s)
	}
}

func (p *CaclParser) Op() (localctx IOpContext) {
	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CaclParserRULE_op)
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
		p.SetState(10)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CaclParserAdd)|(1<<CaclParserSub)|(1<<CaclParserEQ)|(1<<CaclParserCHENG)|(1<<CaclParserCHU))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CaclParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaclParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Int() antlr.TerminalNode {
	return s.GetToken(CaclParserInt, 0)
}

func (s *ExprContext) SetVal() ISetValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetValContext)
}

func (s *ExprContext) NL() antlr.TerminalNode {
	return s.GetToken(CaclParserNL, 0)
}

func (s *ExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) AddVal() IAddValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddValContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaclListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaclListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *CaclParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CaclParserRULE_expr)

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

	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.Match(CaclParserInt)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(14)
			p.SetVal()
		}
		{
			p.SetState(15)
			p.Match(CaclParserNL)
		}
		{
			p.SetState(16)
			p.Expr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(18)
			p.AddVal()
		}

	}

	return localctx
}

// IAddValContext is an interface to support dynamic dispatch.
type IAddValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddValContext differentiates from other interfaces.
	IsAddValContext()
}

type AddValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddValContext() *AddValContext {
	var p = new(AddValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CaclParserRULE_addVal
	return p
}

func (*AddValContext) IsAddValContext() {}

func NewAddValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddValContext {
	var p = new(AddValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaclParserRULE_addVal

	return p
}

func (s *AddValContext) GetParser() antlr.Parser { return s.parser }

func (s *AddValContext) Op() IOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *AddValContext) AllId() []antlr.TerminalNode {
	return s.GetTokens(CaclParserId)
}

func (s *AddValContext) Id(i int) antlr.TerminalNode {
	return s.GetToken(CaclParserId, i)
}

func (s *AddValContext) AllInt() []antlr.TerminalNode {
	return s.GetTokens(CaclParserInt)
}

func (s *AddValContext) Int(i int) antlr.TerminalNode {
	return s.GetToken(CaclParserInt, i)
}

func (s *AddValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaclListener); ok {
		listenerT.EnterAddVal(s)
	}
}

func (s *AddValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaclListener); ok {
		listenerT.ExitAddVal(s)
	}
}

func (p *CaclParser) AddVal() (localctx IAddValContext) {
	localctx = NewAddValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CaclParserRULE_addVal)
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
		p.SetState(21)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CaclParserId || _la == CaclParserInt) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(22)
		p.Op()
	}
	{
		p.SetState(23)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CaclParserId || _la == CaclParserInt) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISetValContext is an interface to support dynamic dispatch.
type ISetValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetValContext differentiates from other interfaces.
	IsSetValContext()
}

type SetValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetValContext() *SetValContext {
	var p = new(SetValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CaclParserRULE_setVal
	return p
}

func (*SetValContext) IsSetValContext() {}

func NewSetValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetValContext {
	var p = new(SetValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaclParserRULE_setVal

	return p
}

func (s *SetValContext) GetParser() antlr.Parser { return s.parser }

func (s *SetValContext) Id() antlr.TerminalNode {
	return s.GetToken(CaclParserId, 0)
}

func (s *SetValContext) EQ() antlr.TerminalNode {
	return s.GetToken(CaclParserEQ, 0)
}

func (s *SetValContext) Int() antlr.TerminalNode {
	return s.GetToken(CaclParserInt, 0)
}

func (s *SetValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaclListener); ok {
		listenerT.EnterSetVal(s)
	}
}

func (s *SetValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaclListener); ok {
		listenerT.ExitSetVal(s)
	}
}

func (p *CaclParser) SetVal() (localctx ISetValContext) {
	localctx = NewSetValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CaclParserRULE_setVal)

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
		p.SetState(25)
		p.Match(CaclParserId)
	}
	{
		p.SetState(26)
		p.Match(CaclParserEQ)
	}
	{
		p.SetState(27)
		p.Match(CaclParserInt)
	}

	return localctx
}

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CaclParserRULE_init
	return p
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaclParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaclListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaclListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *CaclParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CaclParserRULE_init)

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
		p.SetState(29)
		p.Expr()
	}

	return localctx
}
