package main

import (
	"github.com/abusizhishen/calculate/interactive/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	interactive()
}

func calc(input antlr.CharStream)  {
	lexer := parser.NewCalcLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer,antlr.LexerDefaultTokenChannel)
	p := parser.NewCalcParser(tokens)

	antlr.ParseTreeWalkerDefault.Walk(newCalculateListen(),p.Rows())
}

func fileStream()  {
	input,err := antlr.NewFileStream("input.file")
	if err != nil{
		panic(err)
	}

	calc(input)
}

func strStream(s string)  {
	calc(antlr.NewInputStream(s))
}


func interactive()  {
	for {
		strStream(input())
	}
}
