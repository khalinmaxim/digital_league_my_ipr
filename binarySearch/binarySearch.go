package binarySearch

import (
	"fmt"
)

// BinarySearch Бинарный поиск в отсортированном массиве
func BinarySearch(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return true
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

// Start Запуск примера
func Start(arr []int, target int) {
	if BinarySearch(arr, target) {
		fmt.Println("Значение найдено в массиве.")
	} else {
		fmt.Println("Значение не найдено в массиве.")
	}
}
