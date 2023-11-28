package main

import (
	"fmt"
	"sort"
)

/*
pecahannya: 10,25,5,1 --> sort ddari besar ke kecil
uang saya: 42

42: 1,1,1,1,1,1,1 ... 42x = 42 pecahan
42: 5,5,5,5,5,5,5,5,1,1 =42 --> 10 pecahan
42: 10,10,10,10,1,1 --> 6 pecahan
42: 25,5,5,5,1,1 --> 6 pecahan

42: 25,10,5,1,1--> 5 pecahan

100: 25,25,25,25
56; 25,25,5,1
*/

func coinChange(uang int, coinValue []int) []int {
	sort.SliceStable(coinValue, func(i, j int) bool {
		return coinValue[i] > coinValue[j]
	})

	// variabel penampung
	var result []int
	// membaca per koin
	for _, coin := range coinValue {
		// perulangan, ketika uangnya lebih besar dari coin, maka tukar dengan pecahan tsb
		// selama nominal uangnya > coin, maka tetap coin tsb yang diambil
		for uang >= coin {
			// simpan pecahannya
			result = append(result, coin)
			// kurangi nominal uangnya
			uang -= coin
		}
	}

	return result

}

func main() {
	var coinValue = []int{10, 25, 5, 1}
	var money = 56
	fmt.Println(coinChange(money, coinValue))
}
