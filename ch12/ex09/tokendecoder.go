package ex09

import (
	"io"
	"text/scanner"
	"fmt"
	"strconv"
)

type Token interface{}

type Symbol struct {
	Name string
}
type String struct {
	Value string
}
type Int struct {
	Value int
}
type StartList struct {
}
type EndList struct {
}
type lexer struct {
	scan  scanner.Scanner
	token rune
}
type Decoder struct {
	lex *lexer
}

func NewDecoder(reader io.Reader) *Decoder {
	var decoder Decoder
	decoder.lex = &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	decoder.lex.scan.Init(reader)
	decoder.lex.next()
	return &decoder

}
func (d *Decoder) Token() (Token, error) {
	switch d.lex.token {
	case scanner.Ident:
		name := d.lex.text()
		return Symbol{name}, nil
	case scanner.String:
		s, _ := strconv.Unquote(d.lex.text())
		d.lex.next()
		return String{s}, nil
	case scanner.Int:
		i, _ := strconv.Atoi(d.lex.text())
		d.lex.next()
		return Int{i}, nil
	case '(':
		d.lex.next()
		return StartList{}, nil
	case ')':
		d.lex.next()
		return EndList{}, nil
	}
	panic(fmt.Sprintf("unexpected token %q", d.lex.text()))

}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan()
}
func (lex *lexer) text() string {
	return lex.scan.TokenText()
}
func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

