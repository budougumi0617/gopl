// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

func main() {
	x := IntSet{words: []uint64{1, 2, 3}}
	fmt.Fprintf(stdout, "%q", x)
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements
func (s *IntSet) Len() int {
	var result int
	for _, bits := range s.words {
		bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
		bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
		bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
		bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
		r64 := (bits & 0x0000ffff) + (bits >> 16 & 0x0000ffff)
		result += int(r64)
	}
	return result
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for len(s.words) == 0 || word >= len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = make([]uint64, 0, len(s.words))
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	x := new(IntSet)
	for _, e := range s.words {
		x.words = append(x.words, e)
	}
	return x
}

// AddAll allallows a list of values to be added, such as s.AddAll(1, 2, 3).
func (s *IntSet) AddAll(values ...int) {
	for _, v := range values {
		s.Add(v)
	}
}

// IntersectWith sets s to the union intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	var temp = make([]uint64, len(s.words))
	for i, tword := range t.words {
		if i < len(temp) {
			temp[i] = s.words[i] & tword
		}
	}
	s.words = temp
}

// DifferenceWith sets s to the difference union intersection of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= ^tword // remove exist elements in t
		}
	}
}

// SymmetricDifference sets s to the symmetric difference union intersection of s and t
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword // remove exist elements s and t
		} else {
			s.words = append(s.words, tword) // Add only exist elements in t
		}
	}
}
