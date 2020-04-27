// Package expr - ast.go has the following license:
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Adapted by Yiping (Allison) Su
package expr

// Expr represents SQL expression statments
type Expr interface {
	// Eval returns the value of the expression
	Eval(env Env) string
}

// Literal represents the parameter string
type Literal string

// Unary represents a unary expression
//
// e.g. sell:"< 400"
type Unary struct {
	// Operator: < or >
	Op string

	// Expression
	//
	// e.g. 400
	X Expr
}

// Binary represents a SQL query expression
type Binary struct {
	// Op represents a parsable operation (AND, OR)
	Op string

	// X, Y represent the left and right hand side values respectively
	X, Y Expr
}
