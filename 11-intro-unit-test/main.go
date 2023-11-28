package main

import (
	"fakhry/calculator/calculate/nilai"
	kl "fakhry/calculator/kali"
	"fmt"
)

func Penjumlahan(bil1, bil2 int) int {
	return bil1 + bil2
}

func Pengurangan(bil1, bil2 int) int {
	return bil1 - bil2
}

func main() {
	hasil := Penjumlahan(2, 3)
	fmt.Println("hasil:", hasil)
	hasil2 := kl.Perkalian(2, 3)
	fmt.Println("hasil kali:", hasil2)
	hasil3 := nilai.CheckNilaiSiswa(90)
	fmt.Println("hasil nilai:", hasil3)
}
