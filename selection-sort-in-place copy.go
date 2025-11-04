package main

import "fmt"

func selectionSortInPlace(arr []int) {
	n := len(arr)
	for i := 0; i<n-1; i++ {
		minIndex := i;
		for j := i+1; j<n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp 
	}
}

func main() {
	numbers := []int{5, 3, 4, 1, 2}

	selectionSortInPlace(numbers)

	fmt.Println("Ordenado (In Place):", numbers)
}