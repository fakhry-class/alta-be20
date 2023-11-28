package main

import "fmt"

func main() {
	var uang = 100

	if uang >= 10000 {
		fmt.Println("saya bisa membeli bakso")
	} else {
		fmt.Println("uang anda kurang, silakan kumpulkan lebih banyak lagi")
	}

	if uang == 100 {
		fmt.Println("uang anda 100")
	}
}
