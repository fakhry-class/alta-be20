package main

import "fmt"

func linierSearch(elements []int, x int) int {
	//alat bantu untuk membaca per index
	for i := 0; i < len(elements); i++ {
		//membandingkan
		if elements[i] == x {
			return i
		}
	}
	//jika tidak ketemu, fungsi akan mengembalikan -1 (asumsi tidak ada index -1)
	return -1
}

func main() {
	data := []int{4, 2, 6, 9, 1, 5, 10, 15, 12, 14}
	cari := 16
	index := linierSearch(data, cari)
	fmt.Println(index)
	if index < 0 {
		fmt.Println("data tidak ditemukan")
	} else {
		fmt.Println("data ditemukan di index ke: ", index)
	}
}
