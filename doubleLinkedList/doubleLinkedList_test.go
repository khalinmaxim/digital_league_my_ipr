package doubleLinkedList

import (
	"testing"
)

func BenchmarkDoubleLinkedListInsert(b *testing.B) {
	var head *Node
	for i := 0; i < b.N; i++ {
		head = Insert(head, i)
	}
}

func BenchmarkDoubleLinkedListSearch(b *testing.B) {
	var head *Node
	for i := 0; i < 100000; i++ {
		head = Insert(head, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Search(head, i%100000)
	}
}
