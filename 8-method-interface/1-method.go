package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

type Student struct {
	Name string
	NIS  string
}

func fullNameFunc(firstName string, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return
}

func (em Employee) fullName(gelar string) string {
	result := em.FirstName + " " + em.LastName + ", " + gelar
	return result
}

func main() {
	emp1 := Employee{
		FirstName: "Bambang",
		LastName:  "pamungkas",
	}

	emp2 := Employee{
		FirstName: "budi",
		LastName:  "sudarsono",
	}
	fmt.Println(fullNameFunc(emp1.FirstName, emp1.LastName))
	fmt.Println(emp1.fullName("S.Pd"))
	fmt.Println(emp2.fullName("S.Kom"))

}
