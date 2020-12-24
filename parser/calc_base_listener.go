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

// EnterExpr is called when production expr is entered.
func (s *BaseCalcListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseCalcListener) ExitExpr(ctx *ExprContext) {}

// EnterRow is called when production row is entered.
func (s *BaseCalcListener) EnterRow(ctx *RowContext) {}

// ExitRow is called when production row is exited.
func (s *BaseCalcListener) ExitRow(ctx *RowContext) {}

// EnterLast is called when production last is entered.
func (s *BaseCalcListener) EnterLast(ctx *LastContext) {}

// ExitLast is called when production last is exited.
func (s *BaseCalcListener) ExitLast(ctx *LastContext) {}

// EnterFile is called when production file is entered.
func (s *BaseCalcListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseCalcListener) ExitFile(ctx *FileContext) {}
