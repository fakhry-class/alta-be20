package main

import "fmt"

func main() {
	var data = map[int]string{}
	var number = 3
	for i := 1; i <= number; i++ {
		var input string
		fmt.Println("masukan sebuah string ke:", i)
		fmt.Scanln(&input)
		data[i] = input
	}

	fmt.Println(data)
	for i := 1; i <= number; i++ {
		fmt.Println("key:", i, "value:", data[i])
	}
}
