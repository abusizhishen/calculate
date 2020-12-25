// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterNUMBER is called when entering the NUMBER production.
	EnterNUMBER(c *NUMBERContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterCHENGCHU is called when entering the CHENGCHU production.
	EnterCHENGCHU(c *CHENGCHUContext)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// ExitNUMBER is called when exiting the NUMBER production.
	ExitNUMBER(c *NUMBERContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitCHENGCHU is called when exiting the CHENGCHU production.
	ExitCHENGCHU(c *CHENGCHUContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)
}
