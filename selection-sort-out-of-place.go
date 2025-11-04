package main

import "fmt"

func selectionSortOutOfPlace(arr []int) []int {
	n := len(arr)

	result := make([]int, 0, n)
	temp := append([]int(nil), arr...);

	for len(temp) > 0 {
		minIndex := 0
		for i := 1; i<len(temp); i++ {
			if temp[i] < temp[minIndex] {
				minIndex = i
			}
		}
		// adiciona o menor ao resultado
		result = append(result, temp[minIndex])
		// remove o menor da lista temporária
		temp = append(temp[:minIndex], temp[minIndex+1:]...)
	}

	return result;
}

func main() {
	numbers := []int{5, 3, 4, 1, 2}
	sorted := selectionSortOutOfPlace(numbers)
	fmt.Println("Original:", numbers)
	fmt.Println("Ordenado (Out of Place):", sorted)
}