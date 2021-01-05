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
	m[id] = c.pop()
	println(m[id])
}

func (c *CalculateListen)ExitID(ctx  *parser.IDContext)  {
	var id = ctx.Id().GetText()
	if v,ok := m[id];!ok{
		fmt.Println("undefined variable: ", id)
		c.push(0)
	}else {
		c.push(v)
	}


}

func (c *CalculateListen)ExitPOW(ctx  *parser.POWContext)  {
	var fang,nu = c.pop(),c.pop()
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
	var s  []prompt.Suggest
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