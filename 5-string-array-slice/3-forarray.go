package main

import "fmt"

func main() {
	primesNumber := [7]int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println("primes:", primesNumber)
	// fmt.Println("primes0:", primesNumber[0])
	// fmt.Println("primes1:", primesNumber[1])
	// fmt.Println("primes2:", primesNumber[2])
	// fmt.Println("primes3:", primesNumber[3])

	// teknik 1
	for index := 0; index < len(primesNumber); index++ {
		fmt.Println("primes:", primesNumber[index])
	}

	// technique 2
	for index, element := range primesNumber {
		fmt.Println(index, "=>", element)
	}
	for _, value := range primesNumber {
		fmt.Println(value)
	}

	//teknik 3
	index := 0
	for range primesNumber {
		fmt.Println(primesNumber[index])
		index++
	}

}
