package main

import "fmt"

func main() {
	//long
	var nama string
	var umur int8
	fmt.Println("nama anda:", nama)
	fmt.Println("umur anda:", umur)

	var nama1, nama2 string = "andi", "budi"
	fmt.Println("nama anda:", nama1)
	fmt.Println("nama anda:", nama2)
	var nama3 = "rudi"
	nama3 = "tono"
	fmt.Println("nama anda:", nama3)

	var status bool = true
	fmt.Println("stat", status)

	var ipk float32 = 3.9
	fmt.Println("ipk1:", ipk)
	// assign ulang ke sebuah variabels
	ipk = 3.95
	fmt.Println("ipk2:", ipk)

	// short declaration variabel angka1
	angka1 := 0
	// memberikan nilai baru ke angka1 dengan nilai 10
	angka1 = 10
	fmt.Println("angk1", angka1)
	firstname := "agung"
	fmt.Println("name", firstname)

	// variabel nilainya bisa diubah"
	// constanta, nilainya tidak bisa diubah

	const gravitasi float32 = 9.8
	// gravitasi = 11.0
	fmt.Println("gravitasi:", gravitasi)

}
