package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"cc", "ab", "ba", "ca"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	data1 := []float64{10.2, 10.1, 9.3}
	sort.Float64s(data1)
	fmt.Println("data1: ", data1)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	/*
		increasing = ASCENDING (ASC) --> kecil ke besar
		decreasing = DESCENDING (DESC) --> besar ke kecil

		bagaimana mengurutkan dari besar ke kecil ?
	*/
}
