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
