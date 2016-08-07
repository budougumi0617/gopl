// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"os"
	"testing"
)

var input uint64 = 0x1234567890ABCDEF

func Example_main() {
	os.Args = []string{"popcount", "0x1234567890ABCDEF"}
	main()

	// Output:
	// result 32
	// result 32
	// result 32
}

func bench(n int, f func(uint64) int) {
	for j := 0; j < n; j++ {
		f(uint64(j))
	}
}

func benchPopCountByTable(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := range pc {
			pc[j] = pc[j/2] + byte(j&1)
		}
		bench(n, PopCountByTable)
	}
}

func benchPopCountByClear(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		bench(n, PopCountByClear)
	}
}

func benchPopCountByShifting(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		bench(n, PopCountByShifting)
	}
}
func BenchmarkPopCountByTable10(b *testing.B) {
	benchPopCountByTable(b, 10)
}
func BenchmarkPopCountByClear10(b *testing.B) {
	benchPopCountByClear(b, 10)
}
func BenchmarkPopCountByShifting10(b *testing.B) {
	benchPopCountByShifting(b, 10)
}

func BenchmarkPopCountByTable100(b *testing.B) {
	benchPopCountByTable(b, 100)
}
func BenchmarkPopCountByClear100(b *testing.B) {
	benchPopCountByClear(b, 100)
}
func BenchmarkPopCountByShifting100(b *testing.B) {
	benchPopCountByShifting(b, 100)
}

func BenchmarkPopCountByTable1000(b *testing.B) {
	benchPopCountByTable(b, 1000)
}
func BenchmarkPopCountByClear1000(b *testing.B) {
	benchPopCountByClear(b, 1000)
}
func BenchmarkPopCountByShifting1000(b *testing.B) {
	benchPopCountByShifting(b, 1000)
}
func BenchmarkPopCountByTable5000(b *testing.B) {
	benchPopCountByTable(b, 5000)
}
func BenchmarkPopCountByClear5000(b *testing.B) {
	benchPopCountByClear(b, 5000)
}
func BenchmarkPopCountByShifting5000(b *testing.B) {
	benchPopCountByShifting(b, 5000)
}
func BenchmarkPopCountByTable10000(b *testing.B) {
	benchPopCountByTable(b, 10000)
}
func BenchmarkPopCountByClear10000(b *testing.B) {
	benchPopCountByClear(b, 10000)
}
func BenchmarkPopCountByShifting10000(b *testing.B) {
	benchPopCountByShifting(b, 10000)
}

func BenchmarkPopCountByTable20000(b *testing.B) {
	benchPopCountByTable(b, 20000)
}
func BenchmarkPopCountByClear20000(b *testing.B) {
	benchPopCountByClear(b, 20000)
}
func BenchmarkPopCountByShifting20000(b *testing.B) {
	benchPopCountByShifting(b, 20000)
}
