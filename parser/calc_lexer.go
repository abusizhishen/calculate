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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 59, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 7, 4, 28, 10, 4, 12, 4, 14, 4, 31, 11, 4, 3, 5, 3, 5, 5, 5, 35, 10,
	5, 3, 5, 6, 5, 38, 10, 5, 13, 5, 14, 5, 39, 3, 5, 5, 5, 43, 10, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 6, 9, 52, 10, 9, 13, 9, 14, 9, 53,
	3, 9, 3, 9, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 3, 2, 7, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2,
	50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 45, 45, 47, 47, 4,
	2, 44, 44, 49, 49, 2, 63, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5,
	23, 3, 2, 2, 2, 7, 25, 3, 2, 2, 2, 9, 42, 3, 2, 2, 2, 11, 44, 3, 2, 2,
	2, 13, 46, 3, 2, 2, 2, 15, 48, 3, 2, 2, 2, 17, 51, 3, 2, 2, 2, 19, 57,
	3, 2, 2, 2, 21, 22, 7, 42, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7, 43, 2, 2,
	24, 6, 3, 2, 2, 2, 25, 29, 9, 2, 2, 2, 26, 28, 9, 3, 2, 2, 27, 26, 3, 2,
	2, 2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 8,
	3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 32, 34, 9, 4, 2, 2, 33, 35, 7, 48, 2, 2,
	34, 33, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 37, 3, 2, 2, 2, 36, 38, 9,
	4, 2, 2, 37, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39,
	40, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 43, 9, 4, 2, 2, 42, 32, 3, 2, 2,
	2, 42, 41, 3, 2, 2, 2, 43, 10, 3, 2, 2, 2, 44, 45, 9, 5, 2, 2, 45, 12,
	3, 2, 2, 2, 46, 47, 7, 63, 2, 2, 47, 14, 3, 2, 2, 2, 48, 49, 9, 6, 2, 2,
	49, 16, 3, 2, 2, 2, 50, 52, 7, 34, 2, 2, 51, 50, 3, 2, 2, 2, 52, 53, 3,
	2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55,
	56, 8, 9, 2, 2, 56, 18, 3, 2, 2, 2, 57, 58, 7, 12, 2, 2, 58, 20, 3, 2,
	2, 2, 8, 2, 29, 34, 39, 42, 53, 3, 8, 2, 2,
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
	"", "'('", "')'", "", "", "", "'='", "", "", "'\n'",
}

var lexerSymbolicNames = []string{
	"", "", "", "Id", "Int", "AddSub", "EQ", "CHENGCHU", "WS", "NL",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "Id", "Int", "AddSub", "EQ", "CHENGCHU", "WS", "NL",
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
	CalcLexerT__0     = 1
	CalcLexerT__1     = 2
	CalcLexerId       = 3
	CalcLexerInt      = 4
	CalcLexerAddSub   = 5
	CalcLexerEQ       = 6
	CalcLexerCHENGCHU = 7
	CalcLexerWS       = 8
	CalcLexerNL       = 9
)