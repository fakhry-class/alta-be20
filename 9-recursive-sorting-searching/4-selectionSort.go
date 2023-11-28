package main

import "fmt"

func selectionSort(elements []int) []int {
	n := len(elements)
	for k := 0; k < n; k++ {
		minimal := k
		for j := k + 1; j < n; j++ {
			if elements[j] < elements[minimal] {
				minimal = j
			}
		}
		// swap
		elements[k], elements[minimal] = elements[minimal], elements[k]
	}
	return elements
}

func main() {
	data := []int{4, 2, 6, 9, 1, 5, 10, 15, 12, 14}
	result := selectionSort(data)
	fmt.Println("result: ", result)
}
