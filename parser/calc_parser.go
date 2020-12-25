// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 43, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 5,
	2, 15, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 5, 5, 28, 10, 5, 3, 5, 5, 5, 31, 10, 5, 3, 5, 3, 5, 5, 5, 35,
	10, 5, 3, 6, 7, 6, 38, 10, 6, 12, 6, 14, 6, 41, 11, 6, 3, 6, 2, 2, 7, 2,
	4, 6, 8, 10, 2, 3, 3, 2, 5, 6, 2, 43, 2, 14, 3, 2, 2, 2, 4, 16, 3, 2, 2,
	2, 6, 20, 3, 2, 2, 2, 8, 27, 3, 2, 2, 2, 10, 39, 3, 2, 2, 2, 12, 15, 7,
	5, 2, 2, 13, 15, 7, 6, 2, 2, 14, 12, 3, 2, 2, 2, 14, 13, 3, 2, 2, 2, 15,
	3, 3, 2, 2, 2, 16, 17, 7, 5, 2, 2, 17, 18, 7, 11, 2, 2, 18, 19, 7, 6, 2,
	2, 19, 5, 3, 2, 2, 2, 20, 21, 9, 2, 2, 2, 21, 22, 7, 7, 2, 2, 22, 23, 9,
	2, 2, 2, 23, 7, 3, 2, 2, 2, 24, 28, 5, 4, 3, 2, 25, 28, 5, 2, 2, 2, 26,
	28, 5, 6, 4, 2, 27, 24, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 26, 3, 2, 2,
	2, 28, 34, 3, 2, 2, 2, 29, 31, 7, 3, 2, 2, 30, 29, 3, 2, 2, 2, 30, 31,
	3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 35, 7, 4, 2, 2, 33, 35, 7, 2, 2, 3,
	34, 30, 3, 2, 2, 2, 34, 33, 3, 2, 2, 2, 35, 9, 3, 2, 2, 2, 36, 38, 5, 8,
	5, 2, 37, 36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40,
	3, 2, 2, 2, 40, 11, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 7, 14, 27, 30, 34,
	39,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\r'", "'\n'", "", "", "'+'", "'-'", "'*'", "'/'", "'='",
}
var symbolicNames = []string{
	"", "", "", "Id", "NUMBER", "Add", "Sub", "CHENG", "CHU", "EQ", "WS",
}

var ruleNames = []string{
	"singleValue", "setVal", "jia", "row", "start",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CalcParser struct {
	*antlr.BaseParser
}

func NewCalcParser(input antlr.TokenStream) *CalcParser {
	this := new(CalcParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Calc.g4"

	return this
}

// CalcParser tokens.
const (
	CalcParserEOF    = antlr.TokenEOF
	CalcParserT__0   = 1
	CalcParserT__1   = 2
	CalcParserId     = 3
	CalcParserNUMBER = 4
	CalcParserAdd    = 5
	CalcParserSub    = 6
	CalcParserCHENG  = 7
	CalcParserCHU    = 8
	CalcParserEQ     = 9
	CalcParserWS     = 10
)

// CalcParser rules.
const (
	CalcParserRULE_singleValue = 0
	CalcParserRULE_setVal      = 1
	CalcParserRULE_jia         = 2
	CalcParserRULE_row         = 3
	CalcParserRULE_start       = 4
)

// ISingleValueContext is an interface to support dynamic dispatch.
type ISingleValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleValueContext differentiates from other interfaces.
	IsSingleValueContext()
}

type SingleValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleValueContext() *SingleValueContext {
	var p = new(SingleValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_singleValue
	return p
}

func (*SingleValueContext) IsSingleValueContext() {}

func NewSingleValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleValueContext {
	var p = new(SingleValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_singleValue

	return p
}

func (s *SingleValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleValueContext) CopyFrom(ctx *SingleValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SingleValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IDContext struct {
	*SingleValueContext
}

func NewIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IDContext {
	var p = new(IDContext)

	p.SingleValueContext = NewEmptySingleValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleValueContext))

	return p
}

func (s *IDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IDContext) Id() antlr.TerminalNode {
	return s.GetToken(CalcParserId, 0)
}

func (s *IDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterID(s)
	}
}

func (s *IDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitID(s)
	}
}

type OUTContext struct {
	*SingleValueContext
}

func NewOUTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OUTContext {
	var p = new(OUTContext)

	p.SingleValueContext = NewEmptySingleValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleValueContext))

	return p
}

func (s *OUTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OUTContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, 0)
}

func (s *OUTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterOUT(s)
	}
}

func (s *OUTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitOUT(s)
	}
}

func (p *CalcParser) SingleValue() (localctx ISingleValueContext) {
	localctx = NewSingleValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcParserRULE_singleValue)

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

	p.SetState(12)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcParserId:
		localctx = NewIDContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.Match(CalcParserId)
		}

	case CalcParserNUMBER:
		localctx = NewOUTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(11)
			p.Match(CalcParserNUMBER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISetValContext is an interface to support dynamic dispatch.
type ISetValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsSetValContext differentiates from other interfaces.
	IsSetValContext()
}

type SetValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	value  antlr.Token
}

func NewEmptySetValContext() *SetValContext {
	var p = new(SetValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_setVal
	return p
}

func (*SetValContext) IsSetValContext() {}

func NewSetValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetValContext {
	var p = new(SetValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_setVal

	return p
}

func (s *SetValContext) GetParser() antlr.Parser { return s.parser }

func (s *SetValContext) GetId() antlr.Token { return s.id }

func (s *SetValContext) GetValue() antlr.Token { return s.value }

func (s *SetValContext) SetId(v antlr.Token) { s.id = v }

func (s *SetValContext) SetValue(v antlr.Token) { s.value = v }

func (s *SetValContext) EQ() antlr.TerminalNode {
	return s.GetToken(CalcParserEQ, 0)
}

func (s *SetValContext) Id() antlr.TerminalNode {
	return s.GetToken(CalcParserId, 0)
}

func (s *SetValContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, 0)
}

func (s *SetValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterSetVal(s)
	}
}

func (s *SetValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitSetVal(s)
	}
}

func (p *CalcParser) SetVal() (localctx ISetValContext) {
	localctx = NewSetValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalcParserRULE_setVal)

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
		p.SetState(14)

		var _m = p.Match(CalcParserId)

		localctx.(*SetValContext).id = _m
	}
	{
		p.SetState(15)
		p.Match(CalcParserEQ)
	}
	{
		p.SetState(16)

		var _m = p.Match(CalcParserNUMBER)

		localctx.(*SetValContext).value = _m
	}

	return localctx
}

// IJiaContext is an interface to support dynamic dispatch.
type IJiaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJiaContext differentiates from other interfaces.
	IsJiaContext()
}

type JiaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJiaContext() *JiaContext {
	var p = new(JiaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_jia
	return p
}

func (*JiaContext) IsJiaContext() {}

func NewJiaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JiaContext {
	var p = new(JiaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_jia

	return p
}

func (s *JiaContext) GetParser() antlr.Parser { return s.parser }

func (s *JiaContext) Add() antlr.TerminalNode {
	return s.GetToken(CalcParserAdd, 0)
}

func (s *JiaContext) AllId() []antlr.TerminalNode {
	return s.GetTokens(CalcParserId)
}

func (s *JiaContext) Id(i int) antlr.TerminalNode {
	return s.GetToken(CalcParserId, i)
}

func (s *JiaContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(CalcParserNUMBER)
}

func (s *JiaContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, i)
}

func (s *JiaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JiaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JiaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterJia(s)
	}
}

func (s *JiaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitJia(s)
	}
}

func (p *CalcParser) Jia() (localctx IJiaContext) {
	localctx = NewJiaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcParserRULE_jia)
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
		p.SetState(18)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CalcParserId || _la == CalcParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(19)
		p.Match(CalcParserAdd)
	}
	{
		p.SetState(20)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CalcParserId || _la == CalcParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRowContext is an interface to support dynamic dispatch.
type IRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowContext differentiates from other interfaces.
	IsRowContext()
}

type RowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowContext() *RowContext {
	var p = new(RowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_row
	return p
}

func (*RowContext) IsRowContext() {}

func NewRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowContext {
	var p = new(RowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_row

	return p
}

func (s *RowContext) GetParser() antlr.Parser { return s.parser }

func (s *RowContext) SetVal() ISetValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetValContext)
}

func (s *RowContext) SingleValue() ISingleValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleValueContext)
}

func (s *RowContext) Jia() IJiaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJiaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJiaContext)
}

func (s *RowContext) EOF() antlr.TerminalNode {
	return s.GetToken(CalcParserEOF, 0)
}

func (s *RowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterRow(s)
	}
}

func (s *RowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitRow(s)
	}
}

func (p *CalcParser) Row() (localctx IRowContext) {
	localctx = NewRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalcParserRULE_row)
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(22)
			p.SetVal()
		}

	case 2:
		{
			p.SetState(23)
			p.SingleValue()
		}

	case 3:
		{
			p.SetState(24)
			p.Jia()
		}

	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcParserT__0, CalcParserT__1:
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CalcParserT__0 {
			{
				p.SetState(27)
				p.Match(CalcParserT__0)
			}

		}
		{
			p.SetState(30)
			p.Match(CalcParserT__1)
		}

	case CalcParserEOF:
		{
			p.SetState(31)
			p.Match(CalcParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) AllRow() []IRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowContext)(nil)).Elem())
	var tst = make([]IRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowContext)
		}
	}

	return tst
}

func (s *StartContext) Row(i int) IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *CalcParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalcParserRULE_start)
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CalcParserId || _la == CalcParserNUMBER {
		{
			p.SetState(34)
			p.Row()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
