package binaryTree

import (
	"fmt"
)

// Узел бинарного дерева
type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

// Вставка нового узла в бинарное дерево
func Insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{Data: data, Left: nil, Right: nil}
	} else {
		if data <= root.Data {
			root.Left = Insert(root.Left, data)
		} else {
			root.Right = Insert(root.Right, data)
		}
		return root
	}
}

// Поиск значения в бинарном дереве
func Search(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if root.Data == data {
		return true
	} else if data < root.Data {
		return Search(root.Left, data)
	} else {
		return Search(root.Right, data)
	}
}

// Прямой обход бинарного дерева
func PreOrder(root *Node) {
	if root != nil {
		fmt.Print(root.Data, " ")
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}

// Запуск примера
func Start() {
	var root *Node

	// Вставка элементов в бинарное дерево
	root = Insert(root, 50)
	Insert(root, 30)
	Insert(root, 20)
	Insert(root, 40)
	Insert(root, 70)
	Insert(root, 60)
	Insert(root, 80)

	// Поиск значения в бинарном дереве
	if Search(root, 50) {
		fmt.Println("Значение найдено в бинарном дереве.")
	} else {
		fmt.Println("Значение не найдено в бинарном дереве.")
	}

	// Прямой обход бинарного дерева
	fmt.Println("Прямой обход бинарного дерева:")
	PreOrder(root)
	fmt.Println()
}
