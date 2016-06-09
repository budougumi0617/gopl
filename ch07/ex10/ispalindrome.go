package main

import (
	"fmt"
	"sort"
)

type palindrome []rune

func (s palindrome) Len() int           { return len(s) }
func (s palindrome) Less(i, j int) bool { return s[i] != s[j] }
func (s palindrome) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	fmt.Println(IsPalindrome(palindrome([]rune("1234321"))))
	fmt.Println(IsPalindrome(palindrome([]rune("testSet"))))
}

// IsPalindrome returns true if s is palindrome
func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - i - 1
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}
