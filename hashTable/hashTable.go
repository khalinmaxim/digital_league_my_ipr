package hashTable

import (
	"fmt"
)

// Node Узел хэш-таблицы
type Node struct {
	Key   string
	Value int
	Next  *Node
}

// HashTable Хэш-таблица
type HashTable struct {
	Table map[int]*Node
	Size  int
}

// NewHashTable Создание новой хэш-таблицы
func NewHashTable() *HashTable {
	return &HashTable{
		Table: make(map[int]*Node),
		Size:  0,
	}
}

// Хэш-функция
func hashFunction(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash
}

// Insert Вставка значения в хэш-таблицу
func (ht *HashTable) Insert(key string, value int) {
	index := hashFunction(key)
	newNode := &Node{
		Key:   key,
		Value: value,
		Next:  nil,
	}

	if ht.Table[index] == nil {
		ht.Table[index] = newNode
	} else {
		current := ht.Table[index]
		for current.Next != nil {
			if current.Key == key {
				current.Value = value
				return
			}
			current = current.Next
		}
		current.Next = newNode
	}
	ht.Size++
}

// Get Получение значения из хэш-таблицы по ключу
func (ht *HashTable) Get(key string) (int, error) {
	index := hashFunction(key)
	if ht.Table[index] != nil {
		current := ht.Table[index]
		for current != nil {
			if current.Key == key {
				return current.Value, nil
			}
			current = current.Next
		}
	}
	return 0, fmt.Errorf("Key not found")
}

// Start Запуск примера
func Start() {
	// Создание хэш-таблицы
	ht := NewHashTable()

	// Вставка значений в хэш-таблицу
	ht.Insert("apple", 5)
	ht.Insert("banana", 10)
	ht.Insert("orange", 7)
	ht.Insert("grape", 3)

	// Получение значения из хэш-таблицы
	value, err := ht.Get("banana")
	if err == nil {
		fmt.Println("Значение:", value)
	} else {
		fmt.Println(err)
	}
}
