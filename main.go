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

	antlr.ParseTreeWalkerDefault.Walk(newCalcListener(), p.Start())
}

type calcListener struct {
	*parser.BaseCalcListener
	m map[string]int

	stack []int
	idStack []string
}

func newCalcListener() *calcListener {
	return &calcListener{
		BaseCalcListener: new(parser.BaseCalcListener),
		m:                make(map[string]int),
		stack:            nil,
	}
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


//func (l *calcListener) ExitADDSUB(c *parser.ADDSUBContext) {
//	right,left := l.pop(),l.pop()
//
//	switch c.GetOp().GetTokenType() {
//	case parser.CalcLexerAdd:
//		l.push(left+right)
//	case parser.CalcLexerSub:
//		l.push(left-right)
//	default:
//		panic(fmt.Sprintf("unexpected op %s", c.GetOp().GetText()))
//	}
//}

//func (l *calcListener) ExitCHENGCHU(c *parser.CHENGCHUContext) {
//	right,left := l.pop(),l.pop()
//
//	switch c.GetOp().GetTokenType() {
//	case parser.CalcLexerCHENG:
//		l.push(left*right)
//	case parser.CalcLexerCHU:
//		l.push(left/right)
//	default:
//		panic(fmt.Sprintf("unexpected op %s", c.GetOp().GetText()))
//	}
//}

func (l *calcListener) ExitStart(c *parser.StartContext) {
	//log.Printf("result:%d", l.pop())
}

func (l *calcListener) ExitSingleValue(c *parser.SingleValueContext) {

}

func (l *calcListener) EnterSetVal(c *parser.SetValContext) {
	l.m[c.GetId().GetText()] = strToNum(c.GetValue().GetText())
}

func (l *calcListener) ExitSetVal(c *parser.SetValContext) {
	key := c.GetId().GetText()
	log.Printf("%s = %d", key, l.m[key])
}

func strToNum(s string) int {
	i,err := strconv.Atoi(s)
	if err != nil{
		panic(fmt.Sprintf("invalid number:%s", s))
	}

	return i
}

func (l *calcListener) ExitID(c *parser.IDContext) {
	key := c.GetText()
	log.Println(key)
	log.Printf("out: %d", l.m[key])
}

func (l *calcListener) ExitOUT(c *parser.OUTContext) {
	key := c.GetText()
	log.Println(key)
	log.Printf("out: %s", key)
}

