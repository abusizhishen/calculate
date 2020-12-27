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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 56, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 25, 10,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 33, 10, 3, 12, 3, 14, 3, 36,
	11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 44, 10, 4, 3, 5, 3, 5,
	5, 5, 48, 10, 5, 3, 6, 7, 6, 51, 10, 6, 12, 6, 14, 6, 54, 11, 6, 3, 6,
	2, 3, 4, 7, 2, 4, 6, 8, 10, 2, 5, 3, 2, 5, 6, 3, 2, 9, 10, 3, 2, 7, 8,
	2, 60, 2, 12, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 43, 3, 2, 2, 2, 8, 45,
	3, 2, 2, 2, 10, 52, 3, 2, 2, 2, 12, 13, 9, 2, 2, 2, 13, 14, 7, 12, 2, 2,
	14, 15, 9, 2, 2, 2, 15, 3, 3, 2, 2, 2, 16, 17, 8, 3, 1, 2, 17, 25, 7, 6,
	2, 2, 18, 25, 7, 5, 2, 2, 19, 20, 7, 3, 2, 2, 20, 21, 5, 4, 3, 2, 21, 22,
	7, 4, 2, 2, 22, 25, 3, 2, 2, 2, 23, 25, 5, 2, 2, 2, 24, 16, 3, 2, 2, 2,
	24, 18, 3, 2, 2, 2, 24, 19, 3, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 34, 3,
	2, 2, 2, 26, 27, 12, 8, 2, 2, 27, 28, 9, 3, 2, 2, 28, 33, 5, 4, 3, 9, 29,
	30, 12, 7, 2, 2, 30, 31, 9, 4, 2, 2, 31, 33, 5, 4, 3, 8, 32, 26, 3, 2,
	2, 2, 32, 29, 3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35,
	3, 2, 2, 2, 35, 5, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 38, 7, 5, 2, 2,
	38, 39, 7, 11, 2, 2, 39, 44, 7, 6, 2, 2, 40, 44, 7, 5, 2, 2, 41, 44, 7,
	6, 2, 2, 42, 44, 5, 4, 3, 2, 43, 37, 3, 2, 2, 2, 43, 40, 3, 2, 2, 2, 43,
	41, 3, 2, 2, 2, 43, 42, 3, 2, 2, 2, 44, 7, 3, 2, 2, 2, 45, 47, 5, 6, 4,
	2, 46, 48, 7, 14, 2, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 9,
	3, 2, 2, 2, 49, 51, 5, 8, 5, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2,
	52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 11, 3, 2, 2, 2, 54, 52, 3,
	2, 2, 2, 8, 24, 32, 34, 43, 47, 52,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "", "", "'+'", "'-'", "'*'", "'/'", "'='", "'^'",
}
var symbolicNames = []string{
	"", "", "", "Id", "NUMBER", "Add", "Sub", "CHENG", "CHU", "EQ", "POW",
	"WS", "NL",
}

var ruleNames = []string{
	"pow", "expr", "expression", "row", "rows",
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
	CalcParserPOW    = 10
	CalcParserWS     = 11
	CalcParserNL     = 12
)

// CalcParser rules.
const (
	CalcParserRULE_pow        = 0
	CalcParserRULE_expr       = 1
	CalcParserRULE_expression = 2
	CalcParserRULE_row        = 3
	CalcParserRULE_rows       = 4
)

// IPowContext is an interface to support dynamic dispatch.
type IPowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNu returns the nu token.
	GetNu() antlr.Token

	// GetFang returns the fang token.
	GetFang() antlr.Token

	// SetNu sets the nu token.
	SetNu(antlr.Token)

	// SetFang sets the fang token.
	SetFang(antlr.Token)

	// IsPowContext differentiates from other interfaces.
	IsPowContext()
}

type PowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	nu     antlr.Token
	fang   antlr.Token
}

func NewEmptyPowContext() *PowContext {
	var p = new(PowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_pow
	return p
}

func (*PowContext) IsPowContext() {}

func NewPowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowContext {
	var p = new(PowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_pow

	return p
}

func (s *PowContext) GetParser() antlr.Parser { return s.parser }

func (s *PowContext) GetNu() antlr.Token { return s.nu }

func (s *PowContext) GetFang() antlr.Token { return s.fang }

func (s *PowContext) SetNu(v antlr.Token) { s.nu = v }

func (s *PowContext) SetFang(v antlr.Token) { s.fang = v }

func (s *PowContext) POW() antlr.TerminalNode {
	return s.GetToken(CalcParserPOW, 0)
}

func (s *PowContext) AllId() []antlr.TerminalNode {
	return s.GetTokens(CalcParserId)
}

func (s *PowContext) Id(i int) antlr.TerminalNode {
	return s.GetToken(CalcParserId, i)
}

func (s *PowContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(CalcParserNUMBER)
}

func (s *PowContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, i)
}

func (s *PowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterPow(s)
	}
}

func (s *PowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitPow(s)
	}
}

func (p *CalcParser) Pow() (localctx IPowContext) {
	localctx = NewPowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcParserRULE_pow)
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

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PowContext).nu = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == CalcParserId || _la == CalcParserNUMBER) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PowContext).nu = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(11)
		p.Match(CalcParserPOW)
	}
	{
		p.SetState(12)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PowContext).fang = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == CalcParserId || _la == CalcParserNUMBER) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PowContext).fang = _ri
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
	p.RuleIndex = CalcParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NUMBERContext struct {
	*ExprContext
}

func NewNUMBERContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NUMBERContext {
	var p = new(NUMBERContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NUMBERContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NUMBERContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, 0)
}

func (s *NUMBERContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterNUMBER(s)
	}
}

func (s *NUMBERContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitNUMBER(s)
	}
}

type PARENTHESContext struct {
	*ExprContext
}

func NewPARENTHESContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PARENTHESContext {
	var p = new(PARENTHESContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PARENTHESContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PARENTHESContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PARENTHESContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterPARENTHES(s)
	}
}

func (s *PARENTHESContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitPARENTHES(s)
	}
}

type ADDSUBContext struct {
	*ExprContext
	op antlr.Token
}

func NewADDSUBContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ADDSUBContext {
	var p = new(ADDSUBContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ADDSUBContext) GetOp() antlr.Token { return s.op }

func (s *ADDSUBContext) SetOp(v antlr.Token) { s.op = v }

func (s *ADDSUBContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ADDSUBContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ADDSUBContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ADDSUBContext) Add() antlr.TerminalNode {
	return s.GetToken(CalcParserAdd, 0)
}

func (s *ADDSUBContext) Sub() antlr.TerminalNode {
	return s.GetToken(CalcParserSub, 0)
}

func (s *ADDSUBContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterADDSUB(s)
	}
}

func (s *ADDSUBContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitADDSUB(s)
	}
}

type IDContext struct {
	*ExprContext
}

func NewIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IDContext {
	var p = new(IDContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

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

type PPOWContext struct {
	*ExprContext
}

func NewPPOWContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PPOWContext {
	var p = new(PPOWContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PPOWContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PPOWContext) Pow() IPowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPowContext)
}

func (s *PPOWContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterPPOW(s)
	}
}

func (s *PPOWContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitPPOW(s)
	}
}

type CHENGCHUContext struct {
	*ExprContext
	op antlr.Token
}

func NewCHENGCHUContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CHENGCHUContext {
	var p = new(CHENGCHUContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CHENGCHUContext) GetOp() antlr.Token { return s.op }

func (s *CHENGCHUContext) SetOp(v antlr.Token) { s.op = v }

func (s *CHENGCHUContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CHENGCHUContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CHENGCHUContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CHENGCHUContext) CHENG() antlr.TerminalNode {
	return s.GetToken(CalcParserCHENG, 0)
}

func (s *CHENGCHUContext) CHU() antlr.TerminalNode {
	return s.GetToken(CalcParserCHU, 0)
}

func (s *CHENGCHUContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterCHENGCHU(s)
	}
}

func (s *CHENGCHUContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitCHENGCHU(s)
	}
}

func (p *CalcParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *CalcParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CalcParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNUMBERContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(15)
			p.Match(CalcParserNUMBER)
		}

	case 2:
		localctx = NewIDContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(16)
			p.Match(CalcParserId)
		}

	case 3:
		localctx = NewPARENTHESContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)
			p.Match(CalcParserT__0)
		}
		{
			p.SetState(18)
			p.expr(0)
		}
		{
			p.SetState(19)
			p.Match(CalcParserT__1)
		}

	case 4:
		localctx = NewPPOWContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(21)
			p.Pow()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCHENGCHUContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expr)
				p.SetState(24)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(25)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CHENGCHUContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserCHENG || _la == CalcParserCHU) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CHENGCHUContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(26)
					p.expr(7)
				}

			case 2:
				localctx = NewADDSUBContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expr)
				p.SetState(27)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(28)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ADDSUBContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserAdd || _la == CalcParserSub) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ADDSUBContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(29)
					p.expr(6)
				}

			}

		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EXPRContext struct {
	*ExpressionContext
}

func NewEXPRContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EXPRContext {
	var p = new(EXPRContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EXPRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EXPRContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EXPRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterEXPR(s)
	}
}

func (s *EXPRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitEXPR(s)
	}
}

type SetValContext struct {
	*ExpressionContext
}

func NewSetValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetValContext {
	var p = new(SetValContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SetValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetValContext) Id() antlr.TerminalNode {
	return s.GetToken(CalcParserId, 0)
}

func (s *SetValContext) EQ() antlr.TerminalNode {
	return s.GetToken(CalcParserEQ, 0)
}

func (s *SetValContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, 0)
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

type OUTIDContext struct {
	*ExpressionContext
}

func NewOUTIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OUTIDContext {
	var p = new(OUTIDContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OUTIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OUTIDContext) Id() antlr.TerminalNode {
	return s.GetToken(CalcParserId, 0)
}

func (s *OUTIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterOUTID(s)
	}
}

func (s *OUTIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitOUTID(s)
	}
}

type OUTNUMBERContext struct {
	*ExpressionContext
}

func NewOUTNUMBERContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OUTNUMBERContext {
	var p = new(OUTNUMBERContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OUTNUMBERContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OUTNUMBERContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, 0)
}

func (s *OUTNUMBERContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterOUTNUMBER(s)
	}
}

func (s *OUTNUMBERContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitOUTNUMBER(s)
	}
}

func (p *CalcParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcParserRULE_expression)

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

	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSetValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.Match(CalcParserId)
		}
		{
			p.SetState(36)
			p.Match(CalcParserEQ)
		}
		{
			p.SetState(37)
			p.Match(CalcParserNUMBER)
		}

	case 2:
		localctx = NewOUTIDContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Match(CalcParserId)
		}

	case 3:
		localctx = NewOUTNUMBERContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.Match(CalcParserNUMBER)
		}

	case 4:
		localctx = NewEXPRContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.expr(0)
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

func (s *RowContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RowContext) NL() antlr.TerminalNode {
	return s.GetToken(CalcParserNL, 0)
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
	{
		p.SetState(43)
		p.Expression()
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CalcParserNL {
		{
			p.SetState(44)
			p.Match(CalcParserNL)
		}

	}

	return localctx
}

// IRowsContext is an interface to support dynamic dispatch.
type IRowsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowsContext differentiates from other interfaces.
	IsRowsContext()
}

type RowsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowsContext() *RowsContext {
	var p = new(RowsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_rows
	return p
}

func (*RowsContext) IsRowsContext() {}

func NewRowsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowsContext {
	var p = new(RowsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_rows

	return p
}

func (s *RowsContext) GetParser() antlr.Parser { return s.parser }

func (s *RowsContext) AllRow() []IRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowContext)(nil)).Elem())
	var tst = make([]IRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowContext)
		}
	}

	return tst
}

func (s *RowsContext) Row(i int) IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *RowsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterRows(s)
	}
}

func (s *RowsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitRows(s)
	}
}

func (p *CalcParser) Rows() (localctx IRowsContext) {
	localctx = NewRowsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalcParserRULE_rows)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalcParserT__0)|(1<<CalcParserId)|(1<<CalcParserNUMBER))) != 0 {
		{
			p.SetState(47)
			p.Row()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *CalcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalcParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
