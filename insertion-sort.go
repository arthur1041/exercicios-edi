package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		current := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = current
	}
}

func main() {
	arr := []int{8, 2, 4, 3, 7}
	insertionSort(arr)
	fmt.Println("Ordenado:", arr)
}