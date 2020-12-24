package main

import (
	"github.com/abusizhishen/calculate/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
)

func main() {
	input,_ := antlr.NewFileStream("input.file")
	lex := parser.NewCalcLexer(input)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(tokens)

	antlr.ParseTreeWalkerDefault.Walk(&calcListener{}, p.File())
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

func (l *calcListener) ExitRow(c *parser.RowContext) {
}

func (l *calcListener) ExitFile(c *parser.FileContext) {
}

func (l *calcListener) ExitLast(c *parser.LastContext) {
}

func (l *calcListener) ExitExpr(c *parser.ExprContext) {
}

func (l *calcListener) EnterRow(c *parser.RowContext) {
}

func (l *calcListener) EnterFile(c *parser.FileContext) {
	for _,file := range c.GetChildren(){
		log.Println(file.GetPayload())
	}


}


func (l *calcListener) EnterExpr(c *parser.ExprContext) {
	log.Println(c.GetText())
}