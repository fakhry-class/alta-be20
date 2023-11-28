package main

import "fmt"

func main() {
	// var ini untuk menyimpan value 10
	bilangan1 := 10
	var bilangan2 = 15
	// var hasil int
	// hasil = bilangan1 - bilangan2
	var hasil = bilangan1 + bilangan2 //25
	// hasil := bilangan1 + bilangan2
	fmt.Println("Hasil:", hasil)
	hasil = hasil + bilangan1
	fmt.Println("Hasil:", hasil) //35

}
