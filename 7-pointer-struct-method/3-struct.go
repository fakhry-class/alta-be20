package main

import "fmt"

type user struct {
	Nama   string
	Alamat string
	Email  string
	Tinggi uint
}

type mobil struct {
	Warna          string
	Merk           string
	Harga          int
	Tipe           string
	TahunPembelian uint
	BahanBakar     string
}

type pet struct {
	Jenis        string
	Warna        string
	Umur         uint
	JenisKelamin string
}

func main() {
	var namaUser1 string = "fulan"
	var alamatUser1 string = "jakarta"
	var emailUser1 string = "fulan@mail.com"
	var tinggiUser1 uint = 170
	fmt.Println(namaUser1, alamatUser1, emailUser1, tinggiUser1)

	//long declaration
	var user1 user
	user1.Nama = "fulan"
	user1.Alamat = "jakarta"
	user1.Email = "fulan@mail.com"
	user1.Tinggi = 170
	fmt.Println(user1)
	fmt.Println(user1.Nama, "tinggal di:", user1.Alamat)

	var user2 user
	user2.Nama = "budi"
	user2.Alamat = "surabaya"
	fmt.Println(user2)

	//long declaration with value
	var petFakhry = pet{"kucing", "putih", 2, "laki-laki"}
	fmt.Println(petFakhry)

	var petBudi = pet{
		Umur:         1,
		Jenis:        "ayam",
		Warna:        "hitam",
		JenisKelamin: "jantan",
	}
	fmt.Println(petBudi)
	//short declaration
	petFakhry1 := pet{"kucing", "putih", 2, "laki-laki"}
	fmt.Println(petFakhry1)

	user3 := new(user)
	user3.Nama = "fulan"
	user3.Alamat = "jakarta"
	user3.Email = "fulan@mail.com"
	user3.Tinggi = 170
	fmt.Println(user3)

}
