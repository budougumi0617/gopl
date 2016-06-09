package main

import (
	"bytes"
	"os"
	"testing"
)

func TestLimitReader(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{[]string{"dumy"}, "20°C added by Celsius.String()\n"},
		{[]string{"dumy", "-temp", "20°C"}, "20°C added by Celsius.String()\n"},
		{[]string{"dumy", "-temp", "32F"}, "0°C added by Celsius.String()\n"},
	}
	for _, test := range tests {
		os.Args = test.args
		stdout = new(bytes.Buffer) // captured output
		main()
		got := stdout.(*bytes.Buffer).String()
		if got != test.expected {
			t.Errorf("Result = %s, Expected %s", got, test.expected)
		}

	}
}
