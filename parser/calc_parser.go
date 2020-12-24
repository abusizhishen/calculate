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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 48, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 21, 10, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 7, 2, 29, 10, 2, 12, 2, 14, 2, 32, 11, 2, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 5, 7, 5, 41, 10, 5, 12, 5, 14, 5, 44, 11, 5, 3,
	5, 3, 5, 3, 5, 2, 3, 2, 6, 2, 4, 6, 8, 2, 2, 2, 49, 2, 20, 3, 2, 2, 2,
	4, 33, 3, 2, 2, 2, 6, 36, 3, 2, 2, 2, 8, 42, 3, 2, 2, 2, 10, 11, 8, 2,
	1, 2, 11, 12, 7, 3, 2, 2, 12, 13, 5, 2, 2, 2, 13, 14, 7, 4, 2, 2, 14, 21,
	3, 2, 2, 2, 15, 21, 7, 5, 2, 2, 16, 21, 7, 6, 2, 2, 17, 18, 7, 5, 2, 2,
	18, 19, 7, 8, 2, 2, 19, 21, 5, 2, 2, 3, 20, 10, 3, 2, 2, 2, 20, 15, 3,
	2, 2, 2, 20, 16, 3, 2, 2, 2, 20, 17, 3, 2, 2, 2, 21, 30, 3, 2, 2, 2, 22,
	23, 12, 8, 2, 2, 23, 24, 7, 9, 2, 2, 24, 29, 5, 2, 2, 9, 25, 26, 12, 7,
	2, 2, 26, 27, 7, 7, 2, 2, 27, 29, 5, 2, 2, 8, 28, 22, 3, 2, 2, 2, 28, 25,
	3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2,
	31, 3, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 33, 34, 5, 2, 2, 2, 34, 35, 7, 11,
	2, 2, 35, 5, 3, 2, 2, 2, 36, 37, 5, 2, 2, 2, 37, 38, 7, 2, 2, 3, 38, 7,
	3, 2, 2, 2, 39, 41, 5, 4, 3, 2, 40, 39, 3, 2, 2, 2, 41, 44, 3, 2, 2, 2,
	42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 45, 3, 2, 2, 2, 44, 42, 3,
	2, 2, 2, 45, 46, 5, 6, 4, 2, 46, 9, 3, 2, 2, 2, 6, 20, 28, 30, 42,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "", "", "", "'='", "", "", "'\n'",
}
var symbolicNames = []string{
	"", "", "", "Id", "Int", "AddSub", "EQ", "CHENGCHU", "WS", "NL",
}

var ruleNames = []string{
	"expr", "row", "last", "file",
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
	CalcParserEOF      = antlr.TokenEOF
	CalcParserT__0     = 1
	CalcParserT__1     = 2
	CalcParserId       = 3
	CalcParserInt      = 4
	CalcParserAddSub   = 5
	CalcParserEQ       = 6
	CalcParserCHENGCHU = 7
	CalcParserWS       = 8
	CalcParserNL       = 9
)

// CalcParser rules.
const (
	CalcParserRULE_expr = 0
	CalcParserRULE_row  = 1
	CalcParserRULE_last = 2
	CalcParserRULE_file = 3
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

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) Id() antlr.TerminalNode {
	return s.GetToken(CalcParserId, 0)
}

func (s *ExprContext) Int() antlr.TerminalNode {
	return s.GetToken(CalcParserInt, 0)
}

func (s *ExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(CalcParserEQ, 0)
}

func (s *ExprContext) CHENGCHU() antlr.TerminalNode {
	return s.GetToken(CalcParserCHENGCHU, 0)
}

func (s *ExprContext) AddSub() antlr.TerminalNode {
	return s.GetToken(CalcParserAddSub, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitExpr(s)
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
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(9)
			p.Match(CalcParserT__0)
		}
		{
			p.SetState(10)
			p.expr(0)
		}
		{
			p.SetState(11)
			p.Match(CalcParserT__1)
		}

	case 2:
		{
			p.SetState(13)
			p.Match(CalcParserId)
		}

	case 3:
		{
			p.SetState(14)
			p.Match(CalcParserInt)
		}

	case 4:
		{
			p.SetState(15)
			p.Match(CalcParserId)
		}
		{
			p.SetState(16)
			p.Match(CalcParserEQ)
		}
		{
			p.SetState(17)
			p.expr(1)
		}

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
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expr)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(21)
					p.Match(CalcParserCHENGCHU)
				}
				{
					p.SetState(22)
					p.expr(7)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expr)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(24)
					p.Match(CalcParserAddSub)
				}
				{
					p.SetState(25)
					p.expr(6)
				}

			}

		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
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

func (s *RowContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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
	p.EnterRule(localctx, 2, CalcParserRULE_row)

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
		p.SetState(31)
		p.expr(0)
	}
	{
		p.SetState(32)
		p.Match(CalcParserNL)
	}

	return localctx
}

// ILastContext is an interface to support dynamic dispatch.
type ILastContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLastContext differentiates from other interfaces.
	IsLastContext()
}

type LastContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLastContext() *LastContext {
	var p = new(LastContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_last
	return p
}

func (*LastContext) IsLastContext() {}

func NewLastContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LastContext {
	var p = new(LastContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_last

	return p
}

func (s *LastContext) GetParser() antlr.Parser { return s.parser }

func (s *LastContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LastContext) EOF() antlr.TerminalNode {
	return s.GetToken(CalcParserEOF, 0)
}

func (s *LastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LastContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LastContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterLast(s)
	}
}

func (s *LastContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitLast(s)
	}
}

func (p *CalcParser) Last() (localctx ILastContext) {
	localctx = NewLastContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcParserRULE_last)

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
		p.SetState(34)
		p.expr(0)
	}
	{
		p.SetState(35)
		p.Match(CalcParserEOF)
	}

	return localctx
}

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) Last() ILastContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILastContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILastContext)
}

func (s *FileContext) AllRow() []IRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowContext)(nil)).Elem())
	var tst = make([]IRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowContext)
		}
	}

	return tst
}

func (s *FileContext) Row(i int) IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *CalcParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalcParserRULE_file)

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
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(37)
				p.Row()
			}

		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	{
		p.SetState(43)
		p.Last()
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

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
