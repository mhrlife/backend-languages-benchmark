package main

import (
	"gobench/bucket"
	"testing"
)

func BenchmarkSingle(b *testing.B) {
	bc := bucket.NewBucket()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bc.Get(123)
	}
}
