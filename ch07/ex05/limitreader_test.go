// Copyright 2016 budougumi0617 All Rights Reserved.
package limitreader

import (
	"fmt"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	var tests = []struct {
		s string
		p []byte
	}{
		{"LimitReader\ntest", []byte("p")},
	}
	for _, test := range tests {
		lr := LimitReader(strings.NewReader(test.s), int64(len(test.s)))
		for i := 0; i < len(test.s); i++ {
			if lr.Read(test.p); test.p[0] != test.s[i] {
				t.Errorf("Result = %c, Expected %c", test.p[0], test.s[i])
			}
		}
	}
}

func ExampleRead() {
	lr := LimitReader(strings.NewReader("xy"), 2)
	p := []byte("long")
	lr.Read(p)
	fmt.Printf("p = %s\n", p)

	lr = LimitReader(strings.NewReader("return io.EOF"), 0)
	_, err := lr.Read(p)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	// Output:
	// p = xyng
	// EOF
}
