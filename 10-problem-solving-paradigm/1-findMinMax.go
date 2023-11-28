package main

import "fmt"

func findMinMax(input []int) (min int, max int) {
	//asumsi awal
	min = input[0]
	max = input[0]

	// proses membaca 1 per 1 data slice nya
	for i := 0; i < len(input); i++ {
		// fmt.Println(input[i])
		//membandingkan value data ke i dengan nilai min
		if input[i] < min {
			//jika kondisi terpenuhi, artinya ada yang lebih kecil dari min, maka ganti nilai min.
			min = input[i]
		}

		if input[i] > max {
			max = input[i]
		}
	}

	return
}

func main() {
	// findMaxMin([10, 7, 3, 5, 8, 2, 9])
	var data = []int{4, 10, 7, 3, 5, 8, 2, 9}
	min, max := findMinMax(data)
	fmt.Println("min:", min, "max:", max)
}
