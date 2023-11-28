package main

import "fmt"

/*
factorial:
5! = 5 * 4 * 3 *2 *1 = 120
	= 5*4!
4! = 4 * 3 * 2 *1 = 24
	= 4*3!
3! = 3*2*1 = 6
	= 3*2!
2! = 2*1 = 2
1! = 1
*/

func factorialLoop(n int) int {
	var result int = 1
	for i := n; i >= 1; i-- {
		// fmt.Println(i)
		result *= i
	}
	return result
}

func factorial(n int) int {
	if n == 1 { // base case
		return 1
	} else {
		// recurrent relation
		return n * factorial(n-1)
	}
}

func main() {
	// var angka int = 5
	fmt.Println(factorialLoop(5))
	fmt.Println(factorial(5))
}
