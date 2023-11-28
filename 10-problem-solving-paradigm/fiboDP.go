package main

import "fmt"

/*
0,1,1,2,3,5,8....
*/
func fiboDP(number int) int {
	if number <= 1 {
		return number
	}

	fiboI, fiboIMinSatu, fiboIMinDua := -1, 1, 0
	for i := 2; i <= number; i++ {
		fiboI = fiboIMinSatu + fiboIMinDua
		fiboIMinDua = fiboIMinSatu
		fiboIMinSatu = fiboI
	}
	return fiboI
}

func main() {
	fmt.Println(fiboDP(5))
}
