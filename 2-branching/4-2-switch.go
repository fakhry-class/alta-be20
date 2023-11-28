package main

import "fmt"

func main() {
	var today int = 1
	switch {
	case today == 1 && today >= 0:
		fmt.Println("Today is Monday")
		fallthrough
	case today == 2:
		fmt.Println("Today is Tuesday")
		fallthrough
	default:
		fmt.Println("Invalid Date")
	}

}
