package main

import "fmt"

/*
fibonacci
0,1,1,2,3,5,8 ....

fibo(6) = fibo(5)+fibo(4)
	= fibo(4)+fibo(3) + fibo(3)+ fibo(2)
	= fibo(3)+ fibo(2) + fibo(2)+fibo(1) + fibo(2)+fibo(1) + fibo(1)+fibo(0)
*/

func fibonacci(number int) int {
	if number <= 1 {
		return number
	}
	return fibonacci(number-1) + fibonacci(number-2)
}

func main() {
	fmt.Println(fibonacci(6))
}
