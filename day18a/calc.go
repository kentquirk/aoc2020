package main

import (
	"strconv"
)

// TokenType defines one of the token types in our calculator
type TokenType uint8

// These are token type constants
const (
	TypeNumber TokenType = iota
	TypeNegation
	TypeAdd
	TypeSubtract
	TypeMultiply
)

// Token represents one of the tokens that result from our parse
type Token struct {
	T     TokenType
	Value int
}

func (code *Token) String() string {
	switch code.T {
	case TypeNumber:
		return strconv.Itoa(code.Value)
	case TypeAdd:
		return "+"
	case TypeNegation, TypeSubtract:
		return "-"
	case TypeMultiply:
		return "*"
	}
	return ""
}

// Expression is an expression in our grammar
type Expression struct {
	Code []Token
	Top  int
}

// Init sets up an Expression
func (e *Expression) Init(expression string) {
	e.Code = make([]Token, len(expression))
}

// AddOperator adds an operator to the expression
func (e *Expression) AddOperator(operator TokenType) {
	code, top := e.Code, e.Top
	e.Top++
	code[top].T = operator
}

// AddValue adds a value to the expression
func (e *Expression) AddValue(value string) {
	code, top := e.Code, e.Top
	e.Top++
	code[top].Value = 0
	n, _ := strconv.Atoi(value)
	code[top].Value = n
}

// Evaluate processes the Expression and returns a value for it
func (e *Expression) Evaluate() int {
	stack, top := make([]int, len(e.Code)), 0
	for _, code := range e.Code[0:e.Top] {
		switch code.T {
		case TypeNumber:
			stack[top] = code.Value
			top++
			continue
		case TypeNegation:
			stack[top-1] = -stack[top-1]
			continue
		}
		a, b := &stack[top-2], &stack[top-1]
		top--
		switch code.T {
		case TypeAdd:
			*a += *b
		case TypeSubtract:
			*a -= *b
		case TypeMultiply:
			*a *= *b
		}
	}
	return stack[0]
}
