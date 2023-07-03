package binaryTree

import (
	"testing"
)

func BenchmarkBinaryTreeInsert(b *testing.B) {
	var root *Node
	for i := 0; i < 100; i++ {
		root = Insert(root, i)
	}
}

func BenchmarkBinaryTreeSearch(b *testing.B) {
	var root *Node
	for i := 0; i < 100; i++ {
		root = Insert(root, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Search(root, i%100)
	}
}
