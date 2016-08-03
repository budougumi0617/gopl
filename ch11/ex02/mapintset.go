// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
	"sort"
)

// MapIntSet implements by Map.
type MapIntSet struct {
	is map[int]bool
}

// NewMapIntSet returns MapIntSet.
func NewMapIntSet() *MapIntSet {
	return &MapIntSet{map[int]bool{}}
}

// Has includes int value in this.
func (s *MapIntSet) Has(x int) bool {
	return s.is[x]
}

// Add adds value to this.
func (s *MapIntSet) Add(x int) {
	s.is[x] = true
}

// AddAll adds values to this.
func (s *MapIntSet) AddAll(nums ...int) {
	for _, x := range nums {
		s.is[x] = true
	}
}

// UnionWith unites them.
func (s *MapIntSet) UnionWith(t IntSet) {
	for _, x := range t.Elems() {
		s.is[x] = true
	}
}

// Len returns the number of elements
func (s *MapIntSet) Len() int {
	return len(s.is)
}

// Remove removes x from the set
func (s *MapIntSet) Remove(x int) {
	delete(s.is, x)
}

// Clear removes all elements from the set
func (s *MapIntSet) Clear() {
	s.is = make(map[int]bool)
}

// Copy returns a copy of the set
func (s *MapIntSet) Copy() IntSet {
	new := NewMapIntSet()
	for k := range s.is {
		new.is[k] = true
	}
	return new
}

// String returns the set as a string of the form "{1 2 3}".
func (s *MapIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, e := range s.Elems() {
		if e == 0 {
			continue
		}
		if i != 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", e)
	}
	buf.WriteByte('}')
	return buf.String()
}

// IntersectWith sets s to the union intersection of s and t
func (s *MapIntSet) IntersectWith(t IntSet) {
	elems := s.Elems()
	s.Clear()

	for _, se := range elems {
		if t.Has(se) {
			s.Add(se)
		}
	}
}

// DifferenceWith sets s to the difference union intersection of s and t
func (s *MapIntSet) DifferenceWith(t IntSet) {
	elems := s.Elems()
	s.Clear()

	for _, se := range elems {
		if !t.Has(se) {
			s.Add(se)
		}
	}
}

// SymmetricDifference sets s to the symmetric difference union intersection of s and t
func (s *MapIntSet) SymmetricDifference(t IntSet) {
	for _, te := range t.Elems() {
		if s.Has(te) {
			s.Remove(te)
		} else {
			s.Add(te)
		}
	}
}

// Elems returns a slice containing the elements of the set, suitable for iterating over with a range loop
func (s *MapIntSet) Elems() []int {
	ints := make([]int, 0, len(s.is))
	for x := range s.is {
		ints = append(ints, x)
	}
	sort.IntSlice(ints).Sort()
	return ints
}
