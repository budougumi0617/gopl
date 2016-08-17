// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr   string
		env    Env
		result string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		s := expr.String()
		// parsed again,
		reexpr, err := Parse(s)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		// yield an equivalent tree.
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		regot := fmt.Sprintf("%.6g", reexpr.Eval(test.env))

		//		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.result || regot != test.result {
			t.Errorf("\n%s.Eval() in %v = %q,\n%s.Eval() in %v = %q, result %q\n",
				test.expr, test.env, got, reexpr, test.env, regot, test.result)
		}
	}
}
