package linkedList

import (
	"testing"
)

func BenchmarkLinkedListInsert(b *testing.B) {
	var head *Node
	for i := 0; i < b.N; i++ {
		head = Insert(head, i)
	}
}

func BenchmarkLinkedListSearch(b *testing.B) {
	var head *Node
	for i := 0; i < 100000; i++ {
		head = Insert(head, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Search(head, i%100000)
	}
}
