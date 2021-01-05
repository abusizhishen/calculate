// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalcListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNUMBER is called when production NUMBER is entered.
func (s *BaseCalcListener) EnterNUMBER(ctx *NUMBERContext) {}

// ExitNUMBER is called when production NUMBER is exited.
func (s *BaseCalcListener) ExitNUMBER(ctx *NUMBERContext) {}

// EnterPARENTHES is called when production PARENTHES is entered.
func (s *BaseCalcListener) EnterPARENTHES(ctx *PARENTHESContext) {}

// ExitPARENTHES is called when production PARENTHES is exited.
func (s *BaseCalcListener) ExitPARENTHES(ctx *PARENTHESContext) {}

// EnterADDSUB is called when production ADDSUB is entered.
func (s *BaseCalcListener) EnterADDSUB(ctx *ADDSUBContext) {}

// ExitADDSUB is called when production ADDSUB is exited.
func (s *BaseCalcListener) ExitADDSUB(ctx *ADDSUBContext) {}

// EnterPOW is called when production POW is entered.
func (s *BaseCalcListener) EnterPOW(ctx *POWContext) {}

// ExitPOW is called when production POW is exited.
func (s *BaseCalcListener) ExitPOW(ctx *POWContext) {}

// EnterID is called when production ID is entered.
func (s *BaseCalcListener) EnterID(ctx *IDContext) {}

// ExitID is called when production ID is exited.
func (s *BaseCalcListener) ExitID(ctx *IDContext) {}

// EnterCHENGCHU is called when production CHENGCHU is entered.
func (s *BaseCalcListener) EnterCHENGCHU(ctx *CHENGCHUContext) {}

// ExitCHENGCHU is called when production CHENGCHU is exited.
func (s *BaseCalcListener) ExitCHENGCHU(ctx *CHENGCHUContext) {}

// EnterSetVal is called when production SetVal is entered.
func (s *BaseCalcListener) EnterSetVal(ctx *SetValContext) {}

// ExitSetVal is called when production SetVal is exited.
func (s *BaseCalcListener) ExitSetVal(ctx *SetValContext) {}

// EnterOUTID is called when production OUTID is entered.
func (s *BaseCalcListener) EnterOUTID(ctx *OUTIDContext) {}

// ExitOUTID is called when production OUTID is exited.
func (s *BaseCalcListener) ExitOUTID(ctx *OUTIDContext) {}

// EnterOUTNUMBER is called when production OUTNUMBER is entered.
func (s *BaseCalcListener) EnterOUTNUMBER(ctx *OUTNUMBERContext) {}

// ExitOUTNUMBER is called when production OUTNUMBER is exited.
func (s *BaseCalcListener) ExitOUTNUMBER(ctx *OUTNUMBERContext) {}

// EnterEXPR is called when production EXPR is entered.
func (s *BaseCalcListener) EnterEXPR(ctx *EXPRContext) {}

// ExitEXPR is called when production EXPR is exited.
func (s *BaseCalcListener) ExitEXPR(ctx *EXPRContext) {}

// EnterRow is called when production row is entered.
func (s *BaseCalcListener) EnterRow(ctx *RowContext) {}

// ExitRow is called when production row is exited.
func (s *BaseCalcListener) ExitRow(ctx *RowContext) {}

// EnterRows is called when production rows is entered.
func (s *BaseCalcListener) EnterRows(ctx *RowsContext) {}

// ExitRows is called when production rows is exited.
func (s *BaseCalcListener) ExitRows(ctx *RowsContext) {}
