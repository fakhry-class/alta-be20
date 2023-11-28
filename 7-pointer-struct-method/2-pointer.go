package main

import "fmt"

func fillCupValue(cup string) string {
	fmt.Println("data awal:", cup)
	cup = "teh"
	fmt.Println("data akhir:", cup)
	return cup
}

func fillCupReference(cup *string) string {
	fmt.Println("data awal:", cup)
	*cup = "teh"
	fmt.Println("data akhir:", cup)
	return *cup
}

func main() {
	var cup string = "kosong"
	fmt.Println("data awal cup di main:", cup)
	// result := fillCupValue(cup)
	result := fillCupReference(&cup)
	fmt.Println("value result:", result)
	fmt.Println("data akhir cup di main:", cup)
}
