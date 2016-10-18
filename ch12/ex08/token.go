package sexpr

import (
	"fmt"
	"strconv"
	"text/scanner"
)

type Token interface{}
type Symbol struct {
	Value string
}
type String struct {
	Value string
}
type Int struct {
	Value int64
}
type StartList struct{}
type EndList struct{}

type tokenResult struct {
	token Token
	err   error
}

func (d *Decoder) Token() (Token, error) {
	if d.c == nil {
		c := make(chan tokenResult)
		d.c = c
		go func() {
			for {
				readToChan(d.lex, c)
			}
		}()
	}
	result := <-d.c
	return result.token, result.err
}

func readToChan(lex *lexer, c chan<- tokenResult) {
	switch lex.token {
	case scanner.Ident:
		name := lex.text()
		c <- tokenResult{Symbol{Value: name}, nil}
		lex.next()
		return
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		c <- tokenResult{String{Value: s}, nil}
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		c <- tokenResult{Int{Value: int64(i)}, nil}
		lex.next()
		return
	case '(':
		c <- tokenResult{StartList{}, nil}
		lex.next()
		for !endList(lex) {
			readToChan(lex, c)
		}
		c <- tokenResult{EndList{}, nil}
		lex.next() // consume ')'
		return
	}
	// これ多分リークしてる...
	for {
		c <- tokenResult{nil, fmt.Errorf("unexpected token %q", lex.text())}
	}
}
