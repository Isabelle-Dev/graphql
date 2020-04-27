// Package expr - parse.go has the following license:
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Adapted by Yiping (Allison) Su
package expr

import (
	"fmt"
	"strings"
	"text/scanner"
	"unicode"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

// describe returns a string describing the current token, for use in errors.
func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.String:
		return fmt.Sprintf("string %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token)) // any other rune
}

// precedence determines order for operations
func precedence(op string) int {
	switch op {
	case "OR":
		return 2
	case "AND":
		return 1
	}
	return 0
}

// ---- parser ----

// Parse parses the input string as a query expression.
//
//   expr = num                         a literal string, e.g. (hello)
//        | '-' expr                    a unary operator (>, <)
//        | expr '+' expr               a binary operator (AND, OR)
//
func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()
	input = format(input)
	lex := new(lexer)
	lex.scan.IsIdentRune = func(ch rune, i int) bool {
		return ch == '-' || unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '\''
	}
	lex.scan.Init(strings.NewReader(input))

	lex.next() // initial lookahead
	e := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, nil
}

// format will replace all spaces with _ so it's able to be parsed
// by rune
//
// Only exception that will retain spacing is for operators
func format(in string) string {
	ret := strings.ReplaceAll(in, " ", "_")
	ret = strings.ReplaceAll(ret, "_AND_", " AND ")
	ret = strings.ReplaceAll(ret, "_OR_", " OR ")
	return ret
}

func parseExpr(lex *lexer) Expr { return parseBinary(lex, 1) }

// binary = unary ('+' binary)*
// parseBinary stops when it encounters an
// operator of lower precedence than prec1.
func parseBinary(lex *lexer, prec1 int) Expr {
	lhs := parseUnary(lex)
	for prec := precedence(lex.text()); prec >= prec1; prec-- {
		for precedence(lex.text()) == prec {
			op := lex.text()
			lex.next()
			rhs := parseBinary(lex, prec+1)
			lhs = Binary{op, lhs, rhs}
		}
	}
	return lhs
}

// unary = primary
//       | '>' expr
func parseUnary(lex *lexer) Expr {
	if lex.token == '<' || lex.token == '>' {
		op := lex.token
		lex.next()
		return Unary{op, parseUnary(lex)}
	}
	return parsePrimary(lex)
}

// primary = id
//         | '(' expr ')'
func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	case '(':
		lex.next() // consume '('
		e := parseExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", lex.describe())
			panic(lexPanic(msg))
		}
		lex.next() // consume ')'
		return e
	default:
		id := lex.text()
		lex.next()
		return Literal(id)
	}
}
