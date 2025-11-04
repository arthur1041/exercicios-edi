package main

import "fmt"

func countingSort(v []int) []int {
	if len(v) == 0 {
		return v
	}
	max := v[0]
	for _, n := range v {
		if n > max {
			max = n
		}
	}
	count := make([]int, max+1)
	for _, n := range v {
		count[n]++
	}
	index := 0
	for i := 0; i <= max; i++ {
		for count[i] > 0 {
			v[index] = i
			index++
			count[i]--
		}
	}
	return v
}

func main() {
	arr := []int{5, 3, 1, 4, 2}
	countingSort(arr)
	fmt.Println("Ordenado: ", arr)
}