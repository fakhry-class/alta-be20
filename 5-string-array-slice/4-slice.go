package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Obtaining a slice from an array `array`
	// array[low:high]
	var primes = [5]int{2, 3, 5, 7, 11}
	var wilayah = [4]string{"malang", "surabaya", "jogja", "jakarta"}

	// Creating a slice from the array
	var part_primes []int = primes[1:4]
	var part_wilayah []string = wilayah[0:3]

	// part_primes = append(part_primes, 10000)
	// menambah data ke slice akan menambah data ke array juga

	fmt.Println(reflect.ValueOf(part_primes).Kind())
	fmt.Println(part_primes)
	fmt.Println(part_wilayah)
}
