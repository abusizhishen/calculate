// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterPow is called when entering the pow production.
	EnterPow(c *PowContext)

	// EnterNUMBER is called when entering the NUMBER production.
	EnterNUMBER(c *NUMBERContext)

	// EnterPARENTHES is called when entering the PARENTHES production.
	EnterPARENTHES(c *PARENTHESContext)

	// EnterADDSUB is called when entering the ADDSUB production.
	EnterADDSUB(c *ADDSUBContext)

	// EnterID is called when entering the ID production.
	EnterID(c *IDContext)

	// EnterPPOW is called when entering the PPOW production.
	EnterPPOW(c *PPOWContext)

	// EnterCHENGCHU is called when entering the CHENGCHU production.
	EnterCHENGCHU(c *CHENGCHUContext)

	// EnterSetVal is called when entering the SetVal production.
	EnterSetVal(c *SetValContext)

	// EnterOUTID is called when entering the OUTID production.
	EnterOUTID(c *OUTIDContext)

	// EnterOUTNUMBER is called when entering the OUTNUMBER production.
	EnterOUTNUMBER(c *OUTNUMBERContext)

	// EnterEXPR is called when entering the EXPR production.
	EnterEXPR(c *EXPRContext)

	// EnterRow is called when entering the row production.
	EnterRow(c *RowContext)

	// EnterRows is called when entering the rows production.
	EnterRows(c *RowsContext)

	// ExitPow is called when exiting the pow production.
	ExitPow(c *PowContext)

	// ExitNUMBER is called when exiting the NUMBER production.
	ExitNUMBER(c *NUMBERContext)

	// ExitPARENTHES is called when exiting the PARENTHES production.
	ExitPARENTHES(c *PARENTHESContext)

	// ExitADDSUB is called when exiting the ADDSUB production.
	ExitADDSUB(c *ADDSUBContext)

	// ExitID is called when exiting the ID production.
	ExitID(c *IDContext)

	// ExitPPOW is called when exiting the PPOW production.
	ExitPPOW(c *PPOWContext)

	// ExitCHENGCHU is called when exiting the CHENGCHU production.
	ExitCHENGCHU(c *CHENGCHUContext)

	// ExitSetVal is called when exiting the SetVal production.
	ExitSetVal(c *SetValContext)

	// ExitOUTID is called when exiting the OUTID production.
	ExitOUTID(c *OUTIDContext)

	// ExitOUTNUMBER is called when exiting the OUTNUMBER production.
	ExitOUTNUMBER(c *OUTNUMBERContext)

	// ExitEXPR is called when exiting the EXPR production.
	ExitEXPR(c *EXPRContext)

	// ExitRow is called when exiting the row production.
	ExitRow(c *RowContext)

	// ExitRows is called when exiting the rows production.
	ExitRows(c *RowsContext)
}
