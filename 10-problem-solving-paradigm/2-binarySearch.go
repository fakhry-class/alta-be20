package main

import "fmt"

func linierSearch(elements []int, x int) int {
	var counter = 0
	//alat bantu untuk membaca per index
	for i := 0; i < len(elements); i++ {
		counter++
		//membandingkan
		if elements[i] == x {
			fmt.Println("counter:", counter)
			return i
		}
	}
	fmt.Println("counter:", counter)
	//jika tidak ketemu, fungsi akan mengembalikan -1 (asumsi tidak ada index -1)
	return -1
}

func binarySearch(elements []int, x int) int {
	var kiri = 0
	var kanan = len(elements) - 1
	var counter = 0

	for kiri <= kanan {
		counter++
		//mencari index tengah
		var tengah = (kiri + kanan) / 2
		// membandingkan angka yg dicari dengan value di index tengah
		if x < elements[tengah] {
			//jika lebih kecil, maka geser tengah
			kanan = tengah - 1
		} else if x > elements[tengah] {
			kiri = tengah + 1
		} else if x == elements[tengah] {
			fmt.Println("counter:", counter)
			return tengah
		}
	}
	fmt.Println("counter:", counter)
	return -1
}

func main() {
	data := []int{12, 15, 15, 19, 24, 31, 53, 59, 60, 64, 97, 100}
	cari := 97
	// index := linierSearch(data, cari)
	index := binarySearch(data, cari)
	fmt.Println(index)
	if index < 0 {
		fmt.Println("data tidak ditemukan")
	} else {
		fmt.Println("data ditemukan di index ke: ", index)
	}
}
