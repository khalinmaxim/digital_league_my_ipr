package hashTable

import (
	"strconv"
	"testing"
)

func BenchmarkHashTableInsert(b *testing.B) {
	ht := NewHashTable()
	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i)
		value := i
		ht.Insert(key, value)
	}
}

func BenchmarkHashTableGet(b *testing.B) {
	ht := NewHashTable()
	for i := 0; i < 100000; i++ {
		key := "key" + strconv.Itoa(i)
		value := i
		ht.Insert(key, value)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i%100000)
		ht.Get(key)
	}
}
