// Code generated from Cacl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Cacl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCaclListener is a complete listener for a parse tree produced by CaclParser.
type BaseCaclListener struct{}

var _ CaclListener = &BaseCaclListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCaclListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCaclListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCaclListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCaclListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterOp is called when production op is entered.
func (s *BaseCaclListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseCaclListener) ExitOp(ctx *OpContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseCaclListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseCaclListener) ExitExpr(ctx *ExprContext) {}

// EnterAddVal is called when production addVal is entered.
func (s *BaseCaclListener) EnterAddVal(ctx *AddValContext) {}

// ExitAddVal is called when production addVal is exited.
func (s *BaseCaclListener) ExitAddVal(ctx *AddValContext) {}

// EnterSetVal is called when production setVal is entered.
func (s *BaseCaclListener) EnterSetVal(ctx *SetValContext) {}

// ExitSetVal is called when production setVal is exited.
func (s *BaseCaclListener) ExitSetVal(ctx *SetValContext) {}

// EnterInit is called when production init is entered.
func (s *BaseCaclListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseCaclListener) ExitInit(ctx *InitContext) {}
