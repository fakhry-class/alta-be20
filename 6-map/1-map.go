package main

import "fmt"

func main() {
	// long declaration
	var salary = map[string]int{}
	fmt.Println(salary)
	salary["andi"] = 10000
	salary["budi"] = 20000
	fmt.Println(salary)
	fmt.Println(salary["budi"])

	var data1 = map[int]int{}
	data1[10] = 100
	data1[20] = 200
	data1[15] = 150
	fmt.Println(data1)

	var data2 = map[int]string{}
	data2[10] = "sepuluh"
	data2[20] = "dua puluh"
	data2[10] = "seratus"
	fmt.Println(data2)

	// long declaration with value
	var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a)

	// short declaration
	salary_b := map[string]int{}
	fmt.Println(salary_b)
	salary_a1 := map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a1)

	// using make
	var salary_c = make(map[string]int)
	salary_c["doe"] = 7000 // assign value
	fmt.Println(salary_c)
}
