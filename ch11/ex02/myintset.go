// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
)

// An MyIntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type MyIntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *MyIntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *MyIntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll allallows a list of values to be added, such as s.AddAll(1, 2, 3).
func (s *MyIntSet) AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
}

// UnionWith sets s to the union of s and t.
func (s *MyIntSet) UnionWith(t IntSet) {
	if bis, ok := t.(*MyIntSet); ok {
		for i, tword := range bis.words {
			if i < len(s.words) {
				s.words[i] |= tword
			} else {
				s.words = append(s.words, tword)
			}
		}
	} else {
		for _, i := range t.Elems() {
			s.Add(i)
		}
	}
}

// Len returns the number of elements
func (s *MyIntSet) Len() int {
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
func (s *MyIntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for len(s.words) == 0 || word >= len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

// Clear removes all elements from the set
func (s *MyIntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Copy return a copy of the set
func (s *MyIntSet) Copy() IntSet {
	new := &MyIntSet{}
	new.words = make([]uint64, len(s.words))
	copy(new.words, s.words)
	return new
}

// String returns the set as a string of the form "{1 2 3}".
func (s *MyIntSet) String() string {
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

// IntersectWith sets s to the union intersection of s and t
func (s *MyIntSet) IntersectWith(t IntSet) {
	elems := s.Elems()
	s.Clear()

	for _, se := range elems {
		if t.Has(se) {
			s.Add(se)
		}
	}
}

// DifferenceWith sets s to the difference union intersection of s and t
func (s *MyIntSet) DifferenceWith(t IntSet) {
	elems := s.Elems()
	s.Clear()

	for _, se := range elems {
		if !t.Has(se) {
			s.Add(se)
		}
	}
}

// SymmetricDifference sets s to the symmetric difference union intersection of s and t
func (s *MyIntSet) SymmetricDifference(t IntSet) {
	for _, te := range t.Elems() {
		if s.Has(te) {
			s.Remove(te)
		} else {
			s.Add(te)
		}
	}
}

// Elems returns a slice containing the elements of the set, suitable for iterating over with a range loop
func (s *MyIntSet) Elems() (e []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				e = append(e, 64*i+j)
			}
		}
	}
	return
}
