package main

import "fmt"

func main() {
	var myAge = 17

	if myAge == 5 {
		fmt.Println("You too young")
		fmt.Println("terlalu muda")
	}
	if myAge == 17 {
		fmt.Println("So Sweet")
	}
	if myAge >= 17 && myAge < 30 {
		fmt.Println("My Age is between 17 and 30")
	}

	dadAge := 9
	if dadAge < myAge {
		fmt.Println(dadAge)
	}

}
