// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

//!+env

// Env is parameters
type Env map[Var]float64

//!-env

//!+Eval1

// Eval for Var
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// Eval for literal
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

//!-Eval1

//!+Eval2

// Eval for unary
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

// Eval for binary
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

// Eval for call
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	case "min":
		min := c.args[0].Eval(env)
		for i := 1; i < len(c.args); i++ {
			if val := c.args[i].Eval(env); val < min {
				min = val
			}
		}
		return min
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (e extract) Eval(env Env) float64 {
	switch e.fn {
	case "min":
		min := e.args[0].Eval(env)
		for i := 1; i < len(e.args); i++ {
			if val := e.args[i].Eval(env); val < min {
				min = val
			}
		}
		return min
	}
	panic(fmt.Sprintf("unsupported function extract: %s", e.fn))
}

//!-Eval2
