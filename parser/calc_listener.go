// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterID is called when entering the ID production.
	EnterID(c *IDContext)

	// EnterOUT is called when entering the OUT production.
	EnterOUT(c *OUTContext)

	// EnterSetVal is called when entering the setVal production.
	EnterSetVal(c *SetValContext)

	// EnterJia is called when entering the jia production.
	EnterJia(c *JiaContext)

	// EnterRow is called when entering the row production.
	EnterRow(c *RowContext)

	// EnterRows is called when entering the rows production.
	EnterRows(c *RowsContext)

	// ExitID is called when exiting the ID production.
	ExitID(c *IDContext)

	// ExitOUT is called when exiting the OUT production.
	ExitOUT(c *OUTContext)

	// ExitSetVal is called when exiting the setVal production.
	ExitSetVal(c *SetValContext)

	// ExitJia is called when exiting the jia production.
	ExitJia(c *JiaContext)

	// ExitRow is called when exiting the row production.
	ExitRow(c *RowContext)

	// ExitRows is called when exiting the rows production.
	ExitRows(c *RowsContext)
}
