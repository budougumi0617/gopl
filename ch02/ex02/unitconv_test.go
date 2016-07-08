// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMan(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	man()
	got := stdout.(*bytes.Buffer).String()
	expected := "\"unitconv NUMBER...\"Convert temperature, length, or weight.\n"
	if got != expected {
		t.Errorf("Result = %q, Expected %q", got, expected)
	}
}

func TestConvert(t *testing.T) {
	var tests = []struct {
		input    []string
		expected string
	}{
		{strings.Split("10 5", " "), "10°F = -12.222222222222221°C, 10°C = 50°F\n10pond = 4.535970244035199kg, 10kg = 22.046pond\n10f = 3.048780487804878m, 10m = 32.8f\n---------------------------\n5°F = -15°C, 5°C = 41°F\n5pond = 2.2679851220175995kg, 5kg = 11.023pond\n5f = 1.524390243902439m, 5m = 16.4f\n---------------------------\n"},
	}
	for _, test := range tests {
		stdout = new(bytes.Buffer) // captured output
		convert(test.input)
		got := stdout.(*bytes.Buffer).String()
		if got != test.expected {
			t.Errorf("Result = %q, Expected %q", got, test.expected)
		}
	}
}

func TestShow(t *testing.T) {
	var tests = []struct {
		num          float64
		tempResult   string
		weightResult string
		lengthResult string
	}{
		{0, "0°F = -17.77777777777778°C, 0°C = 32°F\n", "0pond = 0kg, 0kg = 0pond\n", "0f = 0m, 0m = 0f\n"},
	}

	for _, test := range tests {
		stdout = new(bytes.Buffer) // captured output
		showTemp(test.num)
		got := stdout.(*bytes.Buffer).String()
		if got != test.tempResult {
			t.Errorf("Result = %q, Expected %q", got, test.tempResult)
		}
	}
	for _, test := range tests {
		stdout = new(bytes.Buffer) // captured output
		showWeight(test.num)
		got := stdout.(*bytes.Buffer).String()
		if got != test.weightResult {
			t.Errorf("Result = %q, Expected %q", got, test.weightResult)
		}
	}
	for _, test := range tests {
		stdout = new(bytes.Buffer) // captured output
		showLength(test.num)
		got := stdout.(*bytes.Buffer).String()
		if got != test.lengthResult {
			t.Errorf("Result = %q, Expected %q", got, test.lengthResult)
		}
	}
}
