package main

import "fmt"

type calculate interface {
	luas() int
	keliling() int
	luasKali(bil int) int
}

type persegi struct {
	sisi int
}

type segitiga struct {
	alas   int
	tinggi int
}

// keliling implements calculate.
func (s segitiga) keliling() int {
	return (s.alas * s.tinggi) / 2
}

// luas implements calculate.
func (s segitiga) luas() int {
	return (s.alas * s.tinggi) / 2
}

// luasKali implements calculate.
func (s segitiga) luasKali(bil int) int {
	return (s.alas * s.tinggi) / 2
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

func (p persegi) keliling() int {
	return 4 * p.sisi
}

func (p persegi) luasKali(bil int) int {
	return 4 * p.sisi
}

func main() {
	var p1 = persegi{
		sisi: 10,
	}

	var IPersegi calculate
	IPersegi = p1

	fmt.Println(IPersegi.luas())
	fmt.Println(IPersegi.keliling())

	var ISegitiga calculate
	ISegitiga = segitiga{
		alas:   5,
		tinggi: 10,
	}
	fmt.Println(ISegitiga.luas())

}
