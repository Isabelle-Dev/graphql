// Package expr - eval.go has the following license:
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Adapted by Yiping (Allison) Su
package expr

import (
	"strings"
)

// Env represents the info of the query type
type Env map[string]string

// Eval (Literal) evaluates singular queries
//
// i.e. the string itself
func (l Literal) Eval(env Env) string {
	ret := strings.ToLower(string(l))
	if index := strings.Index(ret, "'"); index != -1 {
		// embedded single quotes in string must be escaped
		ret = strings.ReplaceAll(ret, "'", "''")
	}
	return addName(addQuote(ret), env)
}

// Eval (Binary) evaluates binary part queries
//
// e.g. (color = 'red') OR (color = 'green')
func (b Binary) Eval(env Env) string {
	switch b.Op {
	case "OR":
		return paren(b.X.Eval(env) + " OR " + b.Y.Eval(env))
	case "AND":
		return paren(b.X.Eval(env) + " AND " + b.Y.Eval(env))
	default:
		// case AND again
		return paren(b.X.Eval(env) + " AND " + b.Y.Eval(env))
	}
}

// Eval (Unary) evaluates unary part queries
//
// e.g. sell < 3000
func (u Unary) Eval(env Env) string {
	switch u.Op {
	case "<":
		str := strings.ReplaceAll(strings.ReplaceAll(u.X.Eval(env), "_", ""), "=", "<")
		if env["name"] == "buy" {
			// we only show valid buy objects
			str += " AND buy > 0"
		}
		return str
	case ">":
		str := strings.ReplaceAll(strings.ReplaceAll(u.X.Eval(env), "_", ""), "=", ">")
		return str
	case "<=":
		str := strings.ReplaceAll(strings.ReplaceAll(u.X.Eval(env), "_", ""), "=", "<=")
		if env["name"] == "buy" {
			// we only show valid buy objects
			str += " AND buy > 0"
		}
		return str
	case ">=":
		str := strings.ReplaceAll(strings.ReplaceAll(u.X.Eval(env), "_", ""), "=", ">=")
		return str
	default:
		return ""
	}
}
