package main

import "fmt"

// function having named return parameter
func perkalian(a, b int) (mul int) {
	mul = a * b
	return
}

func main() {
	m := perkalian(5, 5)
	fmt.Println("5 x 5 = ", m)
}
