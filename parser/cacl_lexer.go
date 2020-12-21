// Code generated from Cacl.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 53, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 7, 2,
	26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 3, 6, 3, 32, 10, 3, 13, 3, 14, 3,
	33, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 2, 2, 12, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 3, 2, 6, 5, 2, 67, 92,
	97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 5,
	2, 12, 12, 15, 15, 34, 34, 2, 54, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 3, 23, 3, 2, 2, 2, 5, 31, 3, 2, 2, 2, 7, 35, 3, 2, 2, 2, 9, 37, 3, 2,
	2, 2, 11, 39, 3, 2, 2, 2, 13, 41, 3, 2, 2, 2, 15, 43, 3, 2, 2, 2, 17, 45,
	3, 2, 2, 2, 19, 47, 3, 2, 2, 2, 21, 51, 3, 2, 2, 2, 23, 27, 9, 2, 2, 2,
	24, 26, 9, 3, 2, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3,
	2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 4, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30,
	32, 9, 4, 2, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31, 3, 2, 2,
	2, 33, 34, 3, 2, 2, 2, 34, 6, 3, 2, 2, 2, 35, 36, 7, 45, 2, 2, 36, 8, 3,
	2, 2, 2, 37, 38, 7, 47, 2, 2, 38, 10, 3, 2, 2, 2, 39, 40, 7, 63, 2, 2,
	40, 12, 3, 2, 2, 2, 41, 42, 7, 44, 2, 2, 42, 14, 3, 2, 2, 2, 43, 44, 7,
	49, 2, 2, 44, 16, 3, 2, 2, 2, 45, 46, 7, 12, 2, 2, 46, 18, 3, 2, 2, 2,
	47, 48, 7, 15, 2, 2, 48, 49, 7, 65, 2, 2, 49, 50, 7, 12, 2, 2, 50, 20,
	3, 2, 2, 2, 51, 52, 9, 5, 2, 2, 52, 22, 3, 2, 2, 2, 5, 2, 27, 33, 2,
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
	"", "", "", "'+'", "'-'", "'='", "'*'", "'/'", "'\n'", "'\r?\n'",
}

var lexerSymbolicNames = []string{
	"", "Id", "Int", "Add", "Sub", "EQ", "CHENG", "CHU", "NL", "HUANHANG",
	"WS",
}

var lexerRuleNames = []string{
	"Id", "Int", "Add", "Sub", "EQ", "CHENG", "CHU", "NL", "HUANHANG", "WS",
}

type CaclLexer struct {
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

func NewCaclLexer(input antlr.CharStream) *CaclLexer {

	l := new(CaclLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Cacl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CaclLexer tokens.
const (
	CaclLexerId       = 1
	CaclLexerInt      = 2
	CaclLexerAdd      = 3
	CaclLexerSub      = 4
	CaclLexerEQ       = 5
	CaclLexerCHENG    = 6
	CaclLexerCHU      = 7
	CaclLexerNL       = 8
	CaclLexerHUANHANG = 9
	CaclLexerWS       = 10
)
