package main

import "testing"

func Example_Main() {
	main()
	// Output:
	// true
	// false
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		s        []rune
		expected bool
	}{
		{[]rune("testset"), true},
		{[]rune("tEstest"), false},
		{[]rune("よのなかねかおかおかねかなのよ"), true},
	}

	for _, test := range tests {
		got := IsPalindrome(palindrome(test.s))
		if got != test.expected {
			t.Errorf("s = %q, got %v want %v", test.s, got, test.expected)
		}
	}
}
