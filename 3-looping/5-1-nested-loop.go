package main

import "fmt"

func main() {
	N := 5
	for baris := 0; baris < N; baris++ {
		for kolom := 0; kolom <= baris; kolom++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
