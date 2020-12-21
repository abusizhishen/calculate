package main

import (
	"fmt"
	"github.com/abusizhishen/calculate/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input,_ := antlr.NewFileStream("input.file")
	lex := parser.NewCaclLexer(input)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewCaclParser(tokens)
	p.BuildParseTrees = true
	fmt.Println(p.Init().GetText())
}
