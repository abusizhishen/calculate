// Code generated from Cacl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Cacl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CaclListener is a complete listener for a parse tree produced by CaclParser.
type CaclListener interface {
	antlr.ParseTreeListener

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterAddVal is called when entering the addVal production.
	EnterAddVal(c *AddValContext)

	// EnterSetVal is called when entering the setVal production.
	EnterSetVal(c *SetValContext)

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitAddVal is called when exiting the addVal production.
	ExitAddVal(c *AddValContext)

	// ExitSetVal is called when exiting the setVal production.
	ExitSetVal(c *SetValContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)
}
