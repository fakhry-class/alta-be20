package main

import "fmt"

type Penulis struct {
	Nama  string
	Email string
}

type Buku struct {
	Judul       string
	TahunTerbit uint
	Penerbit    string
	Author      Penulis
	// Author      []Penulis
}

func main() {
	var penulis1 = Penulis{
		Nama:  "Oda",
		Email: "oda@mail.com",
	}

	var buku1 Buku
	buku1.Judul = "One piece"
	buku1.TahunTerbit = 2000
	buku1.Penerbit = "gramedia"
	buku1.Author = penulis1

	fmt.Println(buku1)
	fmt.Println(buku1.Author.Nama)

}
