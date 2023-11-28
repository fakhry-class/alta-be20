package main

import "fmt"

func main() {
	fmt.Println("Masukkan bilangan yang anda inginkan:")
	var bilangan int
	fmt.Scanln(&bilangan)
	fmt.Println("bil:", bilangan)
	hasil := float32(bilangan) / float32(100) //0
	fmt.Printf("hasil: %v\n", hasil)
	fmt.Println("bil:", bilangan)

}
