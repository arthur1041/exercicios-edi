package main

import "fmt"

func mergeSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	middle := n / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := []int{}

	i := 0 
	j := 0 

	// Compara os elementos e adiciona o menor ao resultado
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Copia o que sobrou da esquerda
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Copia o que sobrou da direita
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func main() {
	arr := []int{9, 4, 3, 6, 3, 2, 5, 7, 1, 8}
	sorted := mergeSort(arr)
	fmt.Println("Ordenado:", sorted)
}
