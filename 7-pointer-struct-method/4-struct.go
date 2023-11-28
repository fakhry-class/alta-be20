package main

import "fmt"

type Person struct {
	Nama      string
	Alamat    string
	Email     string
	Tinggi    uint
	NoTelepon []string
}

func main() {
	var budi Person
	budi.Nama = "budi"
	budi.Alamat = "jakarta"
	budi.Email = "budi@mail.com"
	budi.Tinggi = 180
	budi.NoTelepon = append(budi.NoTelepon, "081234")
	budi.NoTelepon = append(budi.NoTelepon, "081456")
	var telp1 = "09876"
	budi.NoTelepon = append(budi.NoTelepon, telp1)
	fmt.Println(budi)
	fmt.Println(budi.Email)
	fmt.Println(budi.NoTelepon)

	for _, telepon := range budi.NoTelepon {
		fmt.Println("telepon: ", telepon)
	}

	var data []Person
	var person1 Person
	person1.Alamat = "malang"
	person1.Email = "rudi@mail.com"
	person1.Nama = "Rudi"
	person1.NoTelepon = append(person1.NoTelepon, "09871")

	data = append(data, person1)
	data = append(data, budi)
	data = append(data, Person{
		Nama:      "dodi",
		Alamat:    "jogja",
		Email:     "dodi@mail.com",
		Tinggi:    175,
		NoTelepon: []string{"09876"},
	})

	for _, value := range data {
		fmt.Println(value.Nama, "-", value.Email)
	}
}
