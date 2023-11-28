package main

import "fmt"

func main() {
	hour := 10
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}

	if hour <= 24 {
		fmt.Println("hour is valid")
	}
}
