// Copyright 2016 budougumi0617 All Rights Reserved.
package word

import (
	"math/rand"
	"testing"
	"time"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 1
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x10000))
		runes[i] = r
		runes[n-1-i] = r
	}
	m := rng.Intn(n + 1)
	var temp []rune
	temp = append(temp, runes[:m]...)
	temp = append(temp, ' ') // Insert space
	temp = append(temp, runes[m:]...)

	m = rng.Intn(n + 1)
	runes = append(make([]rune, 0), temp[:m]...)
	runes = append(runes, ',') // Insert punctuation
	runes = append(runes, temp[m:]...)
	return string(runes)
}

func TestIsPalindromesByRandomeNoisy(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) returns false, seed is %d", p, seed)
		}
	}
}
