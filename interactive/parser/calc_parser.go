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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 52, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 5, 2, 18, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 7, 2, 29, 10, 2, 12, 2, 14, 2, 32, 11, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 5, 3, 40, 10, 3, 3, 4, 3, 4, 5, 4, 44, 10, 4, 3, 5,
	7, 5, 47, 10, 5, 12, 5, 14, 5, 50, 11, 5, 3, 5, 2, 3, 2, 6, 2, 4, 6, 8,
	2, 4, 3, 2, 9, 10, 3, 2, 7, 8, 2, 57, 2, 17, 3, 2, 2, 2, 4, 39, 3, 2, 2,
	2, 6, 41, 3, 2, 2, 2, 8, 48, 3, 2, 2, 2, 10, 11, 8, 2, 1, 2, 11, 18, 7,
	6, 2, 2, 12, 18, 7, 5, 2, 2, 13, 14, 7, 3, 2, 2, 14, 15, 5, 2, 2, 2, 15,
	16, 7, 4, 2, 2, 16, 18, 3, 2, 2, 2, 17, 10, 3, 2, 2, 2, 17, 12, 3, 2, 2,
	2, 17, 13, 3, 2, 2, 2, 18, 30, 3, 2, 2, 2, 19, 20, 12, 8, 2, 2, 20, 21,
	9, 2, 2, 2, 21, 29, 5, 2, 2, 9, 22, 23, 12, 7, 2, 2, 23, 24, 9, 3, 2, 2,
	24, 29, 5, 2, 2, 8, 25, 26, 12, 3, 2, 2, 26, 27, 7, 12, 2, 2, 27, 29, 5,
	2, 2, 4, 28, 19, 3, 2, 2, 2, 28, 22, 3, 2, 2, 2, 28, 25, 3, 2, 2, 2, 29,
	32, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 3, 3, 2, 2,
	2, 32, 30, 3, 2, 2, 2, 33, 34, 7, 5, 2, 2, 34, 35, 7, 11, 2, 2, 35, 40,
	5, 2, 2, 2, 36, 40, 7, 5, 2, 2, 37, 40, 7, 6, 2, 2, 38, 40, 5, 2, 2, 2,
	39, 33, 3, 2, 2, 2, 39, 36, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 38, 3,
	2, 2, 2, 40, 5, 3, 2, 2, 2, 41, 43, 5, 4, 3, 2, 42, 44, 7, 14, 2, 2, 43,
	42, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 7, 3, 2, 2, 2, 45, 47, 5, 6, 4,
	2, 46, 45, 3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49,
	3, 2, 2, 2, 49, 9, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 8, 17, 28, 30, 39, 43,
	48,
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
	"expr", "expression", "row", "rows",
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
	CalcParserRULE_expr       = 0
	CalcParserRULE_expression = 1
	CalcParserRULE_row        = 2
	CalcParserRULE_rows       = 3
)

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

type POWContext struct {
	*ExprContext
}

func NewPOWContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *POWContext {
	var p = new(POWContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *POWContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *POWContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *POWContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *POWContext) POW() antlr.TerminalNode {
	return s.GetToken(CalcParserPOW, 0)
}

func (s *POWContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterPOW(s)
	}
}

func (s *POWContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitPOW(s)
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
	_startState := 0
	p.EnterRecursionRule(localctx, 0, CalcParserRULE_expr, _p)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcParserNUMBER:
		localctx = NewNUMBERContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(9)
			p.Match(CalcParserNUMBER)
		}

	case CalcParserId:
		localctx = NewIDContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(10)
			p.Match(CalcParserId)
		}

	case CalcParserT__0:
		localctx = NewPARENTHESContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(11)
			p.Match(CalcParserT__0)
		}
		{
			p.SetState(12)
			p.expr(0)
		}
		{
			p.SetState(13)
			p.Match(CalcParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCHENGCHUContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expr)
				p.SetState(17)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(18)

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
					p.SetState(19)
					p.expr(7)
				}

			case 2:
				localctx = NewADDSUBContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expr)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(21)

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
					p.SetState(22)
					p.expr(6)
				}

			case 3:
				localctx = NewPOWContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expr)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(24)
					p.Match(CalcParserPOW)
				}
				{
					p.SetState(25)
					p.expr(2)
				}

			}

		}
		p.SetState(30)
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

func (s *SetValContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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
	p.EnterRule(localctx, 2, CalcParserRULE_expression)

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

	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSetValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Match(CalcParserId)
		}
		{
			p.SetState(32)
			p.Match(CalcParserEQ)
		}
		{
			p.SetState(33)
			p.expr(0)
		}

	case 2:
		localctx = NewOUTIDContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Match(CalcParserId)
		}

	case 3:
		localctx = NewOUTNUMBERContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.Match(CalcParserNUMBER)
		}

	case 4:
		localctx = NewEXPRContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(36)
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
	p.EnterRule(localctx, 4, CalcParserRULE_row)
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
		p.SetState(39)
		p.Expression()
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CalcParserNL {
		{
			p.SetState(40)
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
	p.EnterRule(localctx, 6, CalcParserRULE_rows)
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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalcParserT__0)|(1<<CalcParserId)|(1<<CalcParserNUMBER))) != 0 {
		{
			p.SetState(43)
			p.Row()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *CalcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
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

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
