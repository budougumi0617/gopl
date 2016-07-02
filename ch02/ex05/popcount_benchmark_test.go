// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import "testing"

var input uint64 = 0x1234567890ABCDEF

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(input)
	}
}

func BenchmarkPopCountByClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClear(input)
	}
}
