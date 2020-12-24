// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterRow is called when entering the row production.
	EnterRow(c *RowContext)

	// EnterLast is called when entering the last production.
	EnterLast(c *LastContext)

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitRow is called when exiting the row production.
	ExitRow(c *RowContext)

	// ExitLast is called when exiting the last production.
	ExitLast(c *LastContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)
}
