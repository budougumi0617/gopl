// Copyright 2016 budougumi0617 All Rights Reserved.
package word

import (
	"math/rand"
	"testing"
	"time"
)

func randomNoisyPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 2 // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < n-1; i++ {
		r := rune('A' + rng.Intn('Z'-'A'))
		runes[i] = r
	}
	runes[len(runes)-1] = rune(runes[0] + 1)
	return string(runes)
}

func TestIsPalindromesByRandomeNoisy(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomNoisyPalindrome(rng)
		if IsPalindrome(p) == true {
			t.Errorf("IsPalindrome(%q) returns true", p)
		}
	}
}
