package main

import "fmt"

func main() {
	var datemonth = map[string]int{}
	datemonth["januari"] = 31
	datemonth["februari"] = 28
	datemonth["maret"] = 31
	datemonth["april"] = 30
	datemonth["agustus"] = 0
	fmt.Println(datemonth)
	fmt.Println(datemonth["februari"])
	fmt.Println(datemonth["januari"])
	fmt.Println(datemonth["mei"])
	fmt.Println(datemonth["agustus"])

	value, exist := datemonth["desember"]
	fmt.Println("valuenya:", value, "is exist:", exist)
	// valueMei, existMei := datemonth["mei"]
	// fmt.Println("valuenya:", valueMei, "is exist:", existMei)

	if exist == true {
		fmt.Println("datanya ada cuy")
	} else {
		fmt.Println("datanya tidak ada cuy")
	}

	for key, value := range datemonth {
		fmt.Println("->", key, value)
	}

	var data3 = map[bool]int{}
	data3[true] = 1
	data3[false] = 0
	data3[true] = 2
	fmt.Println("len", len(data3))
	var data4 = map[int]bool{}
	data4[8] = true
	data4[6] = false
	data4[10] = true
	data4[11] = true
	fmt.Println("len", len(data4))
}
