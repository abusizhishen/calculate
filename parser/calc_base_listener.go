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

// EnterID is called when production ID is entered.
func (s *BaseCalcListener) EnterID(ctx *IDContext) {}

// ExitID is called when production ID is exited.
func (s *BaseCalcListener) ExitID(ctx *IDContext) {}

// EnterOUT is called when production OUT is entered.
func (s *BaseCalcListener) EnterOUT(ctx *OUTContext) {}

// ExitOUT is called when production OUT is exited.
func (s *BaseCalcListener) ExitOUT(ctx *OUTContext) {}

// EnterSetVal is called when production setVal is entered.
func (s *BaseCalcListener) EnterSetVal(ctx *SetValContext) {}

// ExitSetVal is called when production setVal is exited.
func (s *BaseCalcListener) ExitSetVal(ctx *SetValContext) {}

// EnterJia is called when production jia is entered.
func (s *BaseCalcListener) EnterJia(ctx *JiaContext) {}

// ExitJia is called when production jia is exited.
func (s *BaseCalcListener) ExitJia(ctx *JiaContext) {}

// EnterRow is called when production row is entered.
func (s *BaseCalcListener) EnterRow(ctx *RowContext) {}

// ExitRow is called when production row is exited.
func (s *BaseCalcListener) ExitRow(ctx *RowContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCalcListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCalcListener) ExitStart(ctx *StartContext) {}
