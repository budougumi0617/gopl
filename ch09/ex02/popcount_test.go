// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import "testing"

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkTable(b *testing.B) {
	bench(b, popCountTable)
}

func BenchmarkShiftMask(b *testing.B) {
	bench(b, popCountShiftMask)
}

func BenchmarkShiftValue(b *testing.B) {
	bench(b, popCountShiftValue)
}

func BenchmarkClearRightmost(b *testing.B) {
	bench(b, popCountClearRightmost)
}

func Example_main() {
	main()
	// Output:
	// Exercise is in test file.
}
