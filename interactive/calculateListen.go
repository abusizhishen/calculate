package main

import (
	"fmt"
	"github.com/abusizhishen/calculate/interactive/parser"
	"github.com/c-bata/go-prompt"
	"math"
	"os"
	"strconv"
)

var m = map[string]int{}

type CalculateListen struct {
	*parser.BaseCalcListener
	stack []int
}

func (c *CalculateListen)pop() int {
	if len(c.stack) == 0 {
		panic("empty stack")
	}

	val := c.stack[len(c.stack)-1]
	c.stack = c.stack[0:len(c.stack)-1]

	return val
}

func (c *CalculateListen)push(i int)  {
	c.stack = append(c.stack, i)
}

func newCalculateListen() *CalculateListen {
	return &CalculateListen{
		BaseCalcListener: new(parser.BaseCalcListener),
	}
}

func (c *CalculateListen)ExitSetVal(ctx  *parser.SetValContext)  {
	id := ctx.Id().GetText()
	val := strToInt(ctx.NUMBER().GetText())
	m[id] = val
	println(val)
}

func (c *CalculateListen)ExitID(ctx  *parser.IDContext)  {
	c.push(m[ctx.Id().GetText()])
}

func (c *CalculateListen)ExitPow(ctx  *parser.PowContext)  {
	var nu,fang int

	switch ctx.GetNu().GetTokenType() {
	case parser.CalcLexerNUMBER:
		nu = strToInt(ctx.GetNu().GetText())
	case parser.CalcLexerId:
		if val,ok := m[ctx.GetNu().GetText()];ok{
			nu = val
		}
	default:
		panic(fmt.Sprintf("unknown variable type of :%s", ctx.GetNu().GetText()))
	}

	switch ctx.GetFang().GetTokenType() {
	case parser.CalcLexerNUMBER:
		fang = strToInt(ctx.GetFang().GetText())
	case parser.CalcLexerId:
		if val,ok := m[ctx.GetFang().GetText()];ok{
			fang = val
		}
	default:
		panic(fmt.Sprintf("unknown variable type of :%s", ctx.GetFang().GetText()))
	}

	c.push(int(math.Pow(float64(nu), float64(fang))))
}

func (c *CalculateListen)ExitOUTID(ctx  *parser.OUTIDContext)  {
	id := ctx.Id().GetText()
	if val,ok := m[id];ok{
		println(val)
		return
	}
	println(fmt.Sprintf("undefinded variable: %s", id))
}

func (c *CalculateListen)ExitRow(ctx  *parser.RowContext)  {
	//println(c.pop())
}

func (c *CalculateListen)ExitADDSUB(ctx  *parser.ADDSUBContext)  {
	right,left := c.pop(),c.pop()
	switch ctx.GetOp().GetTokenType() {
	case parser.CalcLexerAdd:
		c.push(left+right)
	case parser.CalcLexerSub:
		c.push(left-right)
	default:
		panic(fmt.Sprintf("undefinded op: %s", ctx.GetOp().GetText()))
	}
}

func (c *CalculateListen)ExitCHENGCHU(ctx  *parser.CHENGCHUContext)  {
	right,left := c.pop(),c.pop()
	switch ctx.GetOp().GetTokenType() {
	case parser.CalcLexerCHENG:
		c.push(left*right)
	case parser.CalcLexerCHU:
		c.push(left/right)
	default:
		panic(fmt.Sprintf("undefinded op: %s", ctx.GetOp().GetText()))
	}
}

func (c *CalculateListen)ExitNUMBER(ctx  *parser.NUMBERContext)  {
	c.push(strToInt(ctx.NUMBER().GetText()))
}

func (c *CalculateListen)ExitOUTNUMBER(ctx  *parser.OUTNUMBERContext)  {
	println(ctx.GetText())
}

func (c *CalculateListen)ExitEXPR(ctx  *parser.EXPRContext)  {
	println(c.pop())
}

func strToInt(s string) int {
	val,err := strconv.Atoi(s)
	if err != nil{
		panic(err)
	}

	return val
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