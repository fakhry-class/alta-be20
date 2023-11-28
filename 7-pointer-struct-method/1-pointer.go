package main

import "fmt"

func main() {
	var namaUser string = "john doe"
	fmt.Println("value namauser:", namaUser) //john doe
	// mendapatkan alamat memory dari sebuah variabel
	fmt.Println("memory address namauser:", &namaUser)

	// deklarasi variabel pointer
	var pointNama *string = &namaUser
	fmt.Println("value pointNama:", pointNama)
	// membaca value yg tersimpan di sebuah alamat memory
	fmt.Println("value dari alamat pointNama:", *pointNama) //john doe
	*pointNama = "shin tae yong"
	fmt.Println("value dari alamat pointNama:", *pointNama) //shin tae yong
	fmt.Println("value namauser:", namaUser)                //shin tae yong

	var data = []int{1, 2, 3}
	var pointData *[]int = &data
	fmt.Println("add data:", &data[0])
	fmt.Println("add data:", &data[1])
	fmt.Println("add point:", &pointData)

}
