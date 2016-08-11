// Copyright 2016 budougumi0617 All Rights Reserved.
package main

// IntSet includes int velues.
type IntSet interface {
	// Has reports whether the set contains the non-negative value x.
	Has(x int) bool
	// Add adds the non-negative value x to the set.
	Add(x int)
	// AddAll allallows a list of values to be added, such as s.AddAll(1, 2, 3).
	AddAll(nums ...int)
	// Len returns the number of elements
	Len() int
	// Remove removes x from the set
	Remove(x int)
	// Clear removes all elements from the set
	Clear()
	// Copy returns a copy of the set
	Copy() IntSet
	// String returns the set as a string of the form "{1 2 3}".
	String() string
	// UnionWith sets s to the union of s and t.
	UnionWith(t IntSet)
	// IntersectWith sets s to the union intersection of s and t
	IntersectWith(t IntSet)
	// DifferenceWith sets s to the difference union intersection of s and t
	DifferenceWith(t IntSet)
	// SymmetricDifference sets s to the symmetric difference union intersection of s and t
	SymmetricDifference(t IntSet)
	// Elems returns a slice containing the elements of the set, suitable for iterating over with a range loop
	Elems() []int
}
