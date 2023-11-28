package main

import "fmt"

func main() {
	var date int = 25
	if date == 25 {
		fmt.Println("gajian")
	} else if date >= 20 {
		fmt.Println("persiapan gajian")
	} else if date >= 10 && date < 20 {
		fmt.Println("harap bersabar")
	} else {
		fmt.Println("semangat")
	}
}
