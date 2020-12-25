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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 42, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 3, 2, 3, 2, 5, 2, 18, 10, 2, 3, 2, 6, 2, 21, 10, 2, 13, 2, 14, 2, 22,
	3, 2, 5, 2, 26, 10, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 6, 7, 37, 10, 7, 13, 7, 14, 7, 38, 3, 7, 3, 7, 2, 2, 8, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 3, 2, 4, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 45, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2,
	5, 27, 3, 2, 2, 2, 7, 29, 3, 2, 2, 2, 9, 31, 3, 2, 2, 2, 11, 33, 3, 2,
	2, 2, 13, 36, 3, 2, 2, 2, 15, 17, 9, 2, 2, 2, 16, 18, 7, 48, 2, 2, 17,
	16, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18, 20, 3, 2, 2, 2, 19, 21, 9, 2, 2,
	2, 20, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23,
	3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24, 26, 9, 2, 2, 2, 25, 15, 3, 2, 2, 2,
	25, 24, 3, 2, 2, 2, 26, 4, 3, 2, 2, 2, 27, 28, 7, 45, 2, 2, 28, 6, 3, 2,
	2, 2, 29, 30, 7, 47, 2, 2, 30, 8, 3, 2, 2, 2, 31, 32, 7, 44, 2, 2, 32,
	10, 3, 2, 2, 2, 33, 34, 7, 49, 2, 2, 34, 12, 3, 2, 2, 2, 35, 37, 9, 3,
	2, 2, 36, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39,
	3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 41, 8, 7, 2, 2, 41, 14, 3, 2, 2, 2,
	7, 2, 17, 22, 25, 38, 3, 8, 2, 2,
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
	"", "", "'+'", "'-'", "'*'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "NUMBER", "Add", "Sub", "CHENG", "CHU", "WS",
}

var lexerRuleNames = []string{
	"NUMBER", "Add", "Sub", "CHENG", "CHU", "WS",
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
	CalcLexerNUMBER = 1
	CalcLexerAdd    = 2
	CalcLexerSub    = 3
	CalcLexerCHENG  = 4
	CalcLexerCHU    = 5
	CalcLexerWS     = 6
)
