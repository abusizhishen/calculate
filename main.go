package main

import (
	"fmt"
	"github.com/abusizhishen/calculate/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"strconv"
)

func main() {
	input,_ := antlr.NewFileStream("input.file")
	lex := parser.NewCalcLexer(input)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(tokens)

	antlr.ParseTreeWalkerDefault.Walk(&calcListener{}, p.Start())
}

type calcListener struct {
	*parser.BaseCalcListener

	stack []int
}

func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *calcListener) ExitNUMBER(c *parser.NUMBERContext) {
	i,err := strconv.Atoi(c.GetText())
	if err != nil{
		panic(fmt.Sprintf("invalid number %s", c.GetText()))
	}
	
	l.push(i)
}


func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right,left := l.pop(),l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcLexerAdd:
		l.push(left+right)
	case parser.CalcLexerSub:
		l.push(left-right)
	default:
		panic(fmt.Sprintf("unexpected op %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitCHENGCHU(c *parser.CHENGCHUContext) {
	right,left := l.pop(),l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcLexerCHENG:
		l.push(left*right)
	case parser.CalcLexerCHU:
		l.push(left/right)
	default:
		panic(fmt.Sprintf("unexpected op %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitStart(c *parser.StartContext) {
	log.Printf("result:%d", l.pop())
}
