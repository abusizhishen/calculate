// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 64, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 6, 4, 33, 10, 4, 13, 4, 14, 4,
	34, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 12, 6, 12, 52, 10, 12, 13, 12, 14, 12, 53, 3, 12,
	3, 12, 3, 13, 5, 13, 59, 10, 13, 3, 13, 3, 13, 5, 13, 63, 10, 13, 2, 2,
	14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12,
	23, 13, 25, 14, 3, 2, 5, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4,
	2, 11, 11, 34, 34, 2, 67, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 29, 3, 2, 2, 2,
	7, 32, 3, 2, 2, 2, 9, 36, 3, 2, 2, 2, 11, 38, 3, 2, 2, 2, 13, 40, 3, 2,
	2, 2, 15, 42, 3, 2, 2, 2, 17, 44, 3, 2, 2, 2, 19, 46, 3, 2, 2, 2, 21, 48,
	3, 2, 2, 2, 23, 51, 3, 2, 2, 2, 25, 62, 3, 2, 2, 2, 27, 28, 7, 42, 2, 2,
	28, 4, 3, 2, 2, 2, 29, 30, 7, 43, 2, 2, 30, 6, 3, 2, 2, 2, 31, 33, 9, 2,
	2, 2, 32, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35,
	3, 2, 2, 2, 35, 8, 3, 2, 2, 2, 36, 37, 9, 3, 2, 2, 37, 10, 3, 2, 2, 2,
	38, 39, 7, 45, 2, 2, 39, 12, 3, 2, 2, 2, 40, 41, 7, 47, 2, 2, 41, 14, 3,
	2, 2, 2, 42, 43, 7, 44, 2, 2, 43, 16, 3, 2, 2, 2, 44, 45, 7, 49, 2, 2,
	45, 18, 3, 2, 2, 2, 46, 47, 7, 63, 2, 2, 47, 20, 3, 2, 2, 2, 48, 49, 7,
	96, 2, 2, 49, 22, 3, 2, 2, 2, 50, 52, 9, 4, 2, 2, 51, 50, 3, 2, 2, 2, 52,
	53, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 3, 2, 2,
	2, 55, 56, 8, 12, 2, 2, 56, 24, 3, 2, 2, 2, 57, 59, 7, 15, 2, 2, 58, 57,
	3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 63, 7, 12, 2, 2,
	61, 63, 7, 2, 2, 3, 62, 58, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 26, 3,
	2, 2, 2, 7, 2, 34, 53, 58, 62, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "", "", "'+'", "'-'", "'*'", "'/'", "'='", "'^'",
}

var lexerSymbolicNames = []string{
	"", "", "", "Id", "NUMBER", "Add", "Sub", "CHENG", "CHU", "EQ", "POW",
	"WS", "NL",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "Id", "NUMBER", "Add", "Sub", "CHENG", "CHU", "EQ", "POW",
	"WS", "NL",
}

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCalcLexer(input antlr.CharStream) *CalcLexer {

	l := new(CalcLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Calc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerT__0   = 1
	CalcLexerT__1   = 2
	CalcLexerId     = 3
	CalcLexerNUMBER = 4
	CalcLexerAdd    = 5
	CalcLexerSub    = 6
	CalcLexerCHENG  = 7
	CalcLexerCHU    = 8
	CalcLexerEQ     = 9
	CalcLexerPOW    = 10
	CalcLexerWS     = 11
	CalcLexerNL     = 12
)
