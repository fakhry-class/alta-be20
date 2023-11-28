package main

import "fmt"

func describe(i any) {
	fmt.Printf("(%v, %T)\n", i, i)
}
// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

func main() {
	// var data int
	// data = "skdjahlkjd"
	// data = 1234

	// var data interface{}
	var data any
	describe(data)

	data = 42
	describe(data)

	data = "hello"
	describe(data)

	data = true
	describe(data)

	data = []int{1, 2, 3}
	describe(data)
}
