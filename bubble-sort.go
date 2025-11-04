package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)

	for pass := 0; pass < n-1; pass++ {
		fullyOrdered := true

		for i := 0; i < n-1-pass; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				fullyOrdered = false
			}
		}

		if fullyOrdered {
			break
		}
	}
}

func main() {
	arr := []int{5, 3, 1, 4, 2}
	bubbleSort(arr)
	fmt.Println("Ordenado: ", arr)
}