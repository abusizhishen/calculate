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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 51, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 6, 4, 29, 10, 4, 13, 4, 14, 4, 30, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 6, 11, 46, 10,
	11, 13, 11, 14, 11, 47, 3, 11, 3, 11, 2, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 3, 2, 5, 5, 2, 67, 92, 97,
	97, 99, 124, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 2, 52, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 3, 23, 3, 2, 2, 2, 5, 25, 3, 2, 2, 2, 7, 28, 3,
	2, 2, 2, 9, 32, 3, 2, 2, 2, 11, 34, 3, 2, 2, 2, 13, 36, 3, 2, 2, 2, 15,
	38, 3, 2, 2, 2, 17, 40, 3, 2, 2, 2, 19, 42, 3, 2, 2, 2, 21, 45, 3, 2, 2,
	2, 23, 24, 7, 15, 2, 2, 24, 4, 3, 2, 2, 2, 25, 26, 7, 12, 2, 2, 26, 6,
	3, 2, 2, 2, 27, 29, 9, 2, 2, 2, 28, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2,
	30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 8, 3, 2, 2, 2, 32, 33, 9, 3,
	2, 2, 33, 10, 3, 2, 2, 2, 34, 35, 7, 45, 2, 2, 35, 12, 3, 2, 2, 2, 36,
	37, 7, 47, 2, 2, 37, 14, 3, 2, 2, 2, 38, 39, 7, 44, 2, 2, 39, 16, 3, 2,
	2, 2, 40, 41, 7, 49, 2, 2, 41, 18, 3, 2, 2, 2, 42, 43, 7, 63, 2, 2, 43,
	20, 3, 2, 2, 2, 44, 46, 9, 4, 2, 2, 45, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2,
	2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50,
	8, 11, 2, 2, 50, 22, 3, 2, 2, 2, 5, 2, 30, 47, 3, 8, 2, 2,
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
	"", "'\r'", "'\n'", "", "", "'+'", "'-'", "'*'", "'/'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "Id", "NUMBER", "Add", "Sub", "CHENG", "CHU", "EQ", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "Id", "NUMBER", "Add", "Sub", "CHENG", "CHU", "EQ", "WS",
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
	CalcLexerWS     = 10
)
