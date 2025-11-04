package main

import "fmt"

func quickSort(arr []int, start int, end int) {
	if start >= end {
		return
	}

	pivotIndex := partition(arr, start, end)

	quickSort(arr, start, pivotIndex-1)
	quickSort(arr, pivotIndex+1, end) 
}

func partition(arr []int, start int, end int) int {
	pivot := arr[end] 
	pIndex := start  

	// Laço percorre do início até antes do pivô
	for i := start; i < end; i++ {
		if arr[i] <= pivot {
			// troca o elemento atual com o da posição pIndex
			temp := arr[i]
			arr[i] = arr[pIndex]
			arr[pIndex] = temp
			pIndex++
		}
	}

	// Coloca o pivô na posição correta (entre menores e maiores)
	temp := arr[pIndex]
	arr[pIndex] = arr[end]
	arr[end] = temp

	return pIndex
}

func main() {
	arr := []int{9, 4, 3, 6, 3, 2, 8, 7, 1, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Ordenado:", arr)
}
