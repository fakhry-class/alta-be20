package main

import (
	"fmt"
	"strings"
)

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored string ", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int", i)
	case bool:
		fmt.Println("i is boolean")
	default:
		fmt.Println("i stored something else", i)
	}
}

func main() {
	explain("Hello World")
	explain(52)
	explain(true)
	explain([]string{"a", "b"})
}
