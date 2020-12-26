package main

import (
	"fmt"
	"github.com/abusizhishen/calculate/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/c-bata/go-prompt"
	"log"
	"os"
	"strconv"
)

func main() {
	for {
		calc(input())
	}
}

type calcListener struct {
	*parser.BaseCalcListener

	stack []int
	idStack []string
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

func (l *calcListener) ExitSingleValue(c *parser.SingleValueContext) {

}

func (l *calcListener) EnterSetVal(c *parser.SetValContext) {
	val, err := strconv.Atoi(c.NUMBER().GetText())
	if err != nil {
		panic(err)
	}
	m[c.Id().GetText()] = val
	print(val)
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
	log.Printf("out: %d", m[key])
}

func (l *calcListener) ExitOUT(c *parser.OUTContext) {
	key := c.GetText()
	log.Println(key)
	log.Printf("out: %s", key)
}

func (l *calcListener) ExitRow(c *parser.RowContext) {
	key := c.GetText()
	log.Println(key)
	log.Printf("out: %s", key)
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "e", Description: "exit"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func input() string{
	s := prompt.Input("> ", completer)
	switch s {
	case "exit","e","quit","q":
		os.Exit(0)
	default:
		return s
	}

	return ""
}

func calc(str string){
	input := antlr.NewInputStream(str)
	lex := parser.NewCalcLexer(input)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(tokens)
	antlr.ParseTreeWalkerDefault.Walk(calcListener{}, p.Row())
}

var m = make(map[string]int)

func print(i interface{})  {
	fmt.Println(i)
}