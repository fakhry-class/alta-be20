package main

import "fmt"

func main() {
	fmt.Println("masukkan angka:")
	var num int
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Println(i, "- halo alta")
	}
}
