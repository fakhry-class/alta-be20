package main

import (
	"fmt"
	"sort"
)

func main() {
	elements := []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
	value := 33
	findIndex := sort.SearchInts(elements, value)
	fmt.Println("findIndex:", findIndex)
	if value == elements[findIndex] {
		fmt.Println("value found in elemenets")
	} else {
		fmt.Println("value not found in elemenets")
	}
}
