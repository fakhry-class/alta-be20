package main

import "fmt"

func main() {
	fmt.Println("hello")
	sayHelloWorld()
	sapaNama("fakhry")
	data := "alta"
	sapaNama(data)
	salam := greeting(12)
	fmt.Println(salam)
}

// func greeting(hour int) {
// 	if hour < 12 {
// 	  fmt.Println("Selamat Pagi")
// 	} else if hour < 18 {
// 	  fmt.Println("Selamat Sore")
// 	} else {
// 	  fmt.Println("Selamat Malam")
// 	}
//   }

// func with parameter
func greeting(hour int) string {
	var ucapan string
	if hour < 12 {
		ucapan = "Selamat Pagi"
	} else if hour < 18 {
		ucapan = "Selamat Sore"
	} else {
		ucapan = "Selamat Malam"
	}
	return ucapan
}

func sayHelloWorld() {
	fmt.Println("Hello world 1a")
	fmt.Println("Hello world 2a")
}

func sapaNama(name string) {
	fmt.Println("halo", name)
}
