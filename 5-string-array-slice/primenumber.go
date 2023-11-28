package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	// your code here
	var faktor int
	if number < 2 {
		return false
	}

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			faktor++
		}
	}

	if faktor == 2 {
		return true
	}
	return false
}

func primeNumber2(number int) bool {
	if number < 2 {
		return false
	} else if number == 2 {
		return true
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func primeNumber3(number int) bool {
	if number < 2 {
		return false
	} else if number == 2 {
		return true
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber3(13)) // true
	// fmt.Println(primeNumber3(1000000007)) // true
	// fmt.Println(primeNumber3(1500450271)) // true
	// fmt.Println(primeNumber2(1000000000)) // false
	// fmt.Println(primeNumber3(10000000019)) // true
	// fmt.Println(primeNumber(10000000033)) // true
}
