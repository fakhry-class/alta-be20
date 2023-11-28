package main

import "fmt"

func main() {
	// i := 0
	// for i < 5 {
	// 	// fmt.Println("Halo alta")
	// 	fmt.Println("nilai i:", i)
	// 	i++
	// }

	// var status bool = false
	// for !status {
	// 	fmt.Println("halo alta")
	// 	i++
	// 	if i == 5 {
	// 		status = true
	// 	}
	// }

	sum := 1
	for sum < 10 {
		sum += sum // sum = sum+sum
		fmt.Println(sum)
	}

}
