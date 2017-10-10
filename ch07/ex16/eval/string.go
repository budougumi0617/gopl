// Copyright 2017 budougumi0617 All Rights Reserved.

package eval

import (
	"bytes"
	"fmt"
)

// String for Var
func (e Var) String() string {
	return string(e)
}

// String for literal
func (e literal) String() string {
	return fmt.Sprintf("%g", e)
}

func (e unary) String() string {
	return fmt.Sprintf("(%c%s)", e.op, e.x)
}

func (e binary) String() string {
	return fmt.Sprintf("(%s %c %s)", e.x, e.op, e.y)
}

func (e call) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s(", e.fn)
	for i, arg := range e.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		write(buf, arg)
	}
	buf.WriteByte(')')
	return buf.String()
}

func (e extract) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s[", e.fn)
	for i, arg := range e.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		write(buf, arg)
	}
	buf.WriteByte(']')
	return buf.String()
}
