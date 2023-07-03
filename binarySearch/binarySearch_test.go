package binarySearch

import (
	"testing"
)

func BenchmarkBinarySearch(b *testing.B) {
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	for i := 0; i < b.N; i++ {
		target := i % 100
		BinarySearch(arr, target)
	}
}
