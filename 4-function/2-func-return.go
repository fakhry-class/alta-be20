package main

import (
	"fmt"
	"math"
)

// singe return value
func calculateSquare(side int, nama string) int {
	// return side * side
	fmt.Println(nama)
	hasil := side * side
	if hasil > 10 {
		fmt.Println("lebih dari 10")
		return hasil
	}

	fmt.Println("kosong")
	return 0

}

// multiple return value
func calculateCircle(diameter float64) (float64, float64) {
	var keliling = math.Pi * math.Pow(diameter/2, 2)
	var luas = math.Pi * diameter
	// return 2 value
	return keliling, luas
}

func main() {
	var nama = "fakhry"
	var sisi int = 2
	result := calculateSquare(sisi, nama)
	fmt.Println("hasil=", result)

	// var diameter float64 = 12
	// keliling, luas := calculateCircle(diameter)
	// fmt.Printf("Keliling lingkarang: %.2f \n", keliling)
	// fmt.Printf("Luas lingkarang: %.2f \n", luas)

}
