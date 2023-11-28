package main

import "fmt"

func main() {
	var bilangan = []int{2, 3, 5, 7}
	fmt.Println("bil awal:", bilangan)
	var nama = "fakhry"
	fmt.Println("bil awal:", "halo", bilangan, "hello", nama)

	bilangan = append(bilangan, 11)
	bilangan = append(bilangan, 13)
	fmt.Println("bil append:", bilangan)

	var data = []float64{10.1, 11.2, 13.5}
	data = append(data, 15.9)
	fmt.Println(data)

	for i := 0; i < len(bilangan); i++ {
		fmt.Println(bilangan[i])
	}

	var bilangan2 = []int{}
	bilangan2 = append(bilangan2, 100, 200)
	bilangan2 = append(bilangan2, bilangan...)
	fmt.Println("bil2:", bilangan2)

}
