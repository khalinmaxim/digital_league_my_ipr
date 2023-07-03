package linkedList

import (
	"fmt"
)

// Node Узел односвязного списка
type Node struct {
	Data int
	Next *Node
}

// Insert Вставка нового узла в конец списка
func Insert(head *Node, data int) *Node {
	newNode := &Node{Data: data, Next: nil}

	if head == nil {
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

	return head
}

// Search Поиск значения в списке
func Search(head *Node, data int) bool {
	current := head
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}

// Start Запуск примера
func Start() {
	var head *Node

	// Вставка элементов в список
	head = Insert(head, 10)
	head = Insert(head, 20)
	head = Insert(head, 30)
	head = Insert(head, 40)
	head = Insert(head, 50)

	// Поиск значения в списке
	if Search(head, 30) {
		fmt.Println("Значение найдено в списке.")
	} else {
		fmt.Println("Значение не найдено в списке.")
	}
}
