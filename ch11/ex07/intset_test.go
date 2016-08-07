// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"math/rand"
	"testing"
)

func Example_main() {
	main()
	// Output:
	// "{1 2 3}"
}

// newIntSets resturns IntSet pair.
func newIntSets() []IntSet {
	return []IntSet{&MyIntSet{}, NewMapIntSet()}
}

// validateIntSets compares contents of IntSet.
func validateIntSets(t *testing.T, intsets []IntSet) {
	felems, selems := intsets[0].Elems(), intsets[1].Elems()
	for i := 0; i < intsets[0].Len(); i++ {
		if felems[i] != selems[i] {
			t.Errorf("[%d]: felem %d, selem %d", i, felems[i], selems[i])
		}
	}
}

func TestLen(t *testing.T) {
	var tests = []struct {
		words []int
	}{
		{[]int{}},
		{[]int{1, 2, 90}},
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}},
	}
	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		if intsets[0].Len() != intsets[1].Len() {
			t.Errorf("Result = %v, Expected %v", intsets[0].Len(), intsets[1].Len())
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		words   []int
		removes []int
	}{
		{[]int{}, []int{1}},
		{[]int{1, 2, 32}, []int{1, 2, 90, 32}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}, []int{200, 4}},
	}

	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		for _, r := range test.removes {
			intsets[0].Remove(r)
			intsets[1].Remove(r)
		}

		validateIntSets(t, intsets)

	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		words []int
	}{
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}},
	}

	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		intsets[0].Clear()
		intsets[1].Clear()
		validateIntSets(t, intsets)
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		words []int
	}{
		{[]int{}},
		{[]int{1, 2, 90}},
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}},
	}

	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		felems := intsets[0].Copy()
		selems := intsets[1].Copy()
		validateIntSets(t, []IntSet{felems, selems})
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		words []int
	}{
		{[]int{}},
		{[]int{1, 2, 90}},
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}},
	}

	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		fs := intsets[0].String()
		ss := intsets[1].String()
		if fs != ss {
			t.Errorf("Error\n%v\n%v", fs, ss)
		}
	}
}

func TestAddAll(t *testing.T) {
	var tests = []struct {
		words  []int
		values []int
	}{
		{[]int{}, []int{1, 30, 400}},
		{[]int{1, 2, 32}, []int{}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 6, 7, 8}, []int{200}},
	}

	for _, test := range tests {
		intsets := newIntSets()
		for _, w := range test.words {
			intsets[0].Add(w)
			intsets[1].Add(w)
		}
		intsets[0].AddAll(test.values...)
		intsets[1].AddAll(test.values...)

		validateIntSets(t, intsets)

	}
}

func TestIntersectWith(t *testing.T) {
	var tests = []struct {
		xelements []int
		yelements []int
	}{
		{[]int{3}, []int{1, 30, 400}},
		{[]int{1, 2, 32}, []int{1, 2, 32}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 6, 7, 200}, []int{200, 2}},
	}

	for _, test := range tests {
		yintsets := newIntSets()
		xintsets := newIntSets()
		for _, e := range test.xelements {
			xintsets[0].Add(e)
			xintsets[1].Add(e)
		}
		for _, e := range test.yelements {
			yintsets[0].Add(e)
			yintsets[1].Add(e)
		}
		xintsets[0].IntersectWith(yintsets[0])
		xintsets[1].IntersectWith(yintsets[1])
		validateIntSets(t, xintsets)

	}
}

func TestDifferenceWith(t *testing.T) {
	var tests = []struct {
		xelements []int
		yelements []int
	}{
		{[]int{3}, []int{1, 30, 400}},
		{[]int{1, 2, 32}, []int{1, 2, 32}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 6, 7, 200}, []int{200, 2}},
	}

	for _, test := range tests {
		yintsets := newIntSets()
		xintsets := newIntSets()
		for _, e := range test.xelements {
			xintsets[0].Add(e)
			xintsets[1].Add(e)
		}
		for _, e := range test.yelements {
			yintsets[0].Add(e)
			yintsets[1].Add(e)
		}
		xintsets[0].DifferenceWith(yintsets[0])
		xintsets[1].DifferenceWith(yintsets[1])
		validateIntSets(t, xintsets)

	}
}

func TestSymmetricDifference(t *testing.T) {
	var tests = []struct {
		xelements []int
		yelements []int
	}{
		{[]int{3}, []int{1, 30, 400}},
		{[]int{1, 2, 32}, []int{1, 2, 32}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 6, 7, 200}, []int{200, 2, 5}},
	}

	for _, test := range tests {
		yintsets := newIntSets()
		xintsets := newIntSets()
		for _, e := range test.xelements {
			xintsets[0].Add(e)
			xintsets[1].Add(e)
		}
		for _, e := range test.yelements {
			yintsets[0].Add(e)
			yintsets[1].Add(e)
		}
		xintsets[0].SymmetricDifference(yintsets[0])
		xintsets[1].SymmetricDifference(yintsets[1])
		validateIntSets(t, xintsets)

	}
}

func TestUnionWith(t *testing.T) {
	var tests = []struct {
		xelements []int
		yelements []int
	}{
		{[]int{3}, []int{1, 30, 400}},
		{[]int{1, 2, 32}, []int{1, 2, 32}},
		{[]int{1, 2, 90}, []int{2}},
		{[]int{1, 2, 3, 4, 6, 7, 200}, []int{200, 2}},
	}

	for _, test := range tests {
		yintsets := newIntSets()
		xintsets := newIntSets()
		for _, e := range test.xelements {
			xintsets[0].Add(e)
			xintsets[1].Add(e)
		}
		for _, e := range test.yelements {
			yintsets[0].Add(e)
			yintsets[1].Add(e)
		}
		xintsets[0].UnionWith(yintsets[0])
		xintsets[1].UnionWith(yintsets[1])
		validateIntSets(t, xintsets)

	}
}

const max = 20000

func addElement(set IntSet, n int) {
	for i := 0; i < n; i++ {
		set.Add(rand.Intn(max))
	}
}

func benchHas(b *testing.B, set IntSet, n int) {
	addElement(set, n)
	for i := 0; i < b.N; i++ {
		set.Has(rand.Intn(max))
	}
}

func benchAdd(b *testing.B, set IntSet, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			set.Add(rand.Intn(max))
		}
		set.Clear()
	}
}

func benchAddAll(b *testing.B, set IntSet, n int) {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(max)
	}
	for i := 0; i < b.N; i++ {
		set.AddAll(ints...)
		set.Clear()
	}
}

func benchUnionWith(bm *testing.B, a, b IntSet, n int) {
	addElement(a, n)
	addElement(b, n)
	for i := 0; i < bm.N; i++ {
		a.UnionWith(b)
	}
}

func benchDifferenceWith(bm *testing.B, a, b IntSet, n int) {
	addElement(a, n)
	addElement(b, n)
	for i := 0; i < bm.N; i++ {
		a.DifferenceWith(b)
	}
}

func benchIntersectWith(bm *testing.B, a, b IntSet, n int) {
	addElement(a, n)
	addElement(b, n)
	for i := 0; i < bm.N; i++ {
		a.IntersectWith(b)
	}
}

func benchSymmetricDifference(bm *testing.B, a, b IntSet, n int) {
	addElement(a, n)
	addElement(b, n)
	for i := 0; i < bm.N; i++ {
		a.SymmetricDifference(b)
	}
}

// Benchmark to Has
func BenchmarkMapIntSetHas10(b *testing.B) {
	benchHas(b, NewMapIntSet(), 10)
}
func BenchmarkMyIntSetHas10(b *testing.B) {
	benchHas(b, NewMyIntSet(), 10)
}
func BenchmarkMapIntSetHas100(b *testing.B) {
	benchHas(b, NewMapIntSet(), 100)
}
func BenchmarkMyIntSetHas100(b *testing.B) {
	benchHas(b, NewMyIntSet(), 100)
}
func BenchmarkMapIntSetHas1000(b *testing.B) {
	benchHas(b, NewMapIntSet(), 1000)
}
func BenchmarkMyIntSetHas1000(b *testing.B) {
	benchHas(b, NewMyIntSet(), 1000)
}

// Benchmark to Add
func BenchmarkMapIntSetAdd10(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 10)
}
func BenchmarkMyIntSetAdd10(b *testing.B) {
	benchAdd(b, NewMyIntSet(), 10)
}
func BenchmarkMapIntSetAdd100(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 100)
}
func BenchmarkMyIntSetAdd100(b *testing.B) {
	benchAdd(b, NewMyIntSet(), 100)
}
func BenchmarkMapIntSetAdd1000(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 1000)
}
func BenchmarkMyIntSetAdd1000(b *testing.B) {
	benchAdd(b, NewMyIntSet(), 1000)
}

// Benchmark to AddAll
func BenchmarkMapIntSetAddAll10(b *testing.B) {
	benchAddAll(b, NewMapIntSet(), 10)
}
func BenchmarkMyIntSetAddAll10(b *testing.B) {
	benchAddAll(b, NewMyIntSet(), 10)
}
func BenchmarkMapIntSetAddAll100(b *testing.B) {
	benchAddAll(b, NewMapIntSet(), 100)
}
func BenchmarkMyIntSetAddAll100(b *testing.B) {
	benchAddAll(b, NewMyIntSet(), 100)
}
func BenchmarkMapIntSetAddAll1000(b *testing.B) {
	benchAddAll(b, NewMapIntSet(), 1000)
}
func BenchmarkMyIntSetAddAll1000(b *testing.B) {
	benchAddAll(b, NewMyIntSet(), 1000)
}

// Benchmark to UnionWith
func BenchmarkMapIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 10)
}
func BenchmarkMyIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewMyIntSet(), NewMyIntSet(), 10)
}
func BenchmarkMapIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 100)
}
func BenchmarkMyIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewMyIntSet(), NewMyIntSet(), 100)
}
func BenchmarkMapIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 1000)
}
func BenchmarkMyIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewMyIntSet(), NewMyIntSet(), 1000)
}

// Benchmark to IntersectWith
func BenchmarkMapIntSetIntersectWith10(b *testing.B) {
	benchIntersectWith(b, NewMapIntSet(), NewMapIntSet(), 10)
}
func BenchmarkMyIntSetIntersectWith10(b *testing.B) {
	benchIntersectWith(b, NewMyIntSet(), NewMyIntSet(), 10)
}
func BenchmarkMapIntSetIntersectWith100(b *testing.B) {
	benchIntersectWith(b, NewMapIntSet(), NewMapIntSet(), 100)
}
func BenchmarkMyIntSetIntersectWith100(b *testing.B) {
	benchIntersectWith(b, NewMyIntSet(), NewMyIntSet(), 100)
}
func BenchmarkMapIntSetIntersectWith1000(b *testing.B) {
	benchIntersectWith(b, NewMapIntSet(), NewMapIntSet(), 1000)
}
func BenchmarkMyIntSetIntersectWith1000(b *testing.B) {
	benchIntersectWith(b, NewMyIntSet(), NewMyIntSet(), 1000)
}

// Benchmark to DifferenceWith
func BenchmarkMapIntSetDifferenceWith10(b *testing.B) {
	benchDifferenceWith(b, NewMapIntSet(), NewMapIntSet(), 10)
}
func BenchmarkMyIntSetDifferenceWith10(b *testing.B) {
	benchDifferenceWith(b, NewMyIntSet(), NewMyIntSet(), 10)
}
func BenchmarkMapIntSetDifferenceWith100(b *testing.B) {
	benchDifferenceWith(b, NewMapIntSet(), NewMapIntSet(), 100)
}
func BenchmarkMyIntSetDifferenceWith100(b *testing.B) {
	benchDifferenceWith(b, NewMyIntSet(), NewMyIntSet(), 100)
}
func BenchmarkMapIntSetDifferenceWith1000(b *testing.B) {
	benchDifferenceWith(b, NewMapIntSet(), NewMapIntSet(), 1000)
}
func BenchmarkMyIntSetDifferenceWith1000(b *testing.B) {
	benchDifferenceWith(b, NewMyIntSet(), NewMyIntSet(), 1000)
}

// Benchmark to SymmetricDifference
func BenchmarkMapIntSetSymmetricDifference10(b *testing.B) {
	benchSymmetricDifference(b, NewMapIntSet(), NewMapIntSet(), 10)
}
func BenchmarkMyIntSetSymmetricDifference10(b *testing.B) {
	benchSymmetricDifference(b, NewMyIntSet(), NewMyIntSet(), 10)
}
func BenchmarkMapIntSetSymmetricDifference100(b *testing.B) {
	benchSymmetricDifference(b, NewMapIntSet(), NewMapIntSet(), 100)
}
func BenchmarkMyIntSetSymmetricDifference100(b *testing.B) {
	benchSymmetricDifference(b, NewMyIntSet(), NewMyIntSet(), 100)
}
func BenchmarkMapIntSetSymmetricDifference1000(b *testing.B) {
	benchSymmetricDifference(b, NewMapIntSet(), NewMapIntSet(), 1000)
}
func BenchmarkMyIntSetSymmetricDifference1000(b *testing.B) {
	benchSymmetricDifference(b, NewMyIntSet(), NewMyIntSet(), 1000)
}
