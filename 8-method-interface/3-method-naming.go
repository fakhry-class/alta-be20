package main

import (
	"fmt"
	"math"
)

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Luas() float64 {
	return r.width * r.height
}

func (c Circle) Luas() float64 {
	return math.Pi * c.radius * c.radius
}

func LuasPersegi(panjang float64, lebar float64) {

}

func LuasLingkaran(radius float64) {

}

func main() {
	rect := Rect{5.0, 4.0}
	cir := Circle{5.0}
	fmt.Printf("Area of rectangle rect = %0.2f\n", rect.Luas())
	fmt.Printf("Area of circle cir = %0.2f\n", cir.Luas())
	// fmt.Printf("Area of circle cir = %0.2f\n", LuasLingkaran(cir.radius))
}
