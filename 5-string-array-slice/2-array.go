package main

import (
	"fmt"
)

func main() {
	var countries [5]string
	countries[0] = "indonesia"
	countries[1] = "malaysia"
	countries[2] = "singapore"
	countries[1] = "inggris"

	fmt.Println(countries)
	fmt.Println(countries[0])

	numbersA := [4]int{0, 1, 2}
	fmt.Println("numbers A:", numbersA)
	var even_numbers [5]int = [5]int{2, 4, 6, 8}
	fmt.Println("even numbers:", even_numbers)

	primes := [...]int{2, 3, 3}
	fmt.Println(primes)
	fmt.Println("panjang primes", len(primes))

	numbersGenap := [5]int{1: 2, 2: 4}
	fmt.Println("numbers genap:", numbersGenap)
}
