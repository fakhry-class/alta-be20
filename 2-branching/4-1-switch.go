package main

import "fmt"

func main() {
	// var today int = 2
	// switch today {
	// case 1:
	// 	fmt.Printf("Today is Monday")
	// case 2:
	// 	fmt.Printf("Today is Tuesday")
	// default: //else
	// 	fmt.Printf("Invalid Date")
	// }
	var today string = "jumat"
	switch today {
	case "senin":
		fmt.Printf("Today is Monday")
	case "selasa":
		fmt.Printf("Today is Tuesday")
	case "jumat":
		fmt.Printf("Today is friday")
	default: //else
		fmt.Printf("Invalid Date")
	}

	// if today == "senin" {
	// 	fmt.Printf("Today is Monday")
	// } else if today == "selasa" {
	// 	fmt.Printf("Today is Tuesday")
	// } else if today == "jumat" {
	// 	fmt.Printf("Today is friday")
	// } else {
	// 	fmt.Printf("Invalid Date")
	// }

}
