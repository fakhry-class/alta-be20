package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// A Response struct to map the Entire Response
type StarWarsPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Films     []string `json:"films"`
	Vehicles  []string `json:"vehicles"`
}

type PeopleResponse struct {
	Count   int              `json:"count"`
	Next    string           `json:"next"`
	Results []StarWarsPeople `json:"results"`
}

// func main() {
// 	response, err := http.Get("https://swapi.dev/api/people/2")
// 	if err != nil {
// 		log.Fatal("error consume API")
// 	}

// 	responseData, errBody := io.ReadAll(response.Body)
// 	if errBody != nil {
// 		log.Fatal("error read response body")
// 	}
// 	defer response.Body.Close()

// 	var people1 StarWarsPeople
// 	json.Unmarshal(responseData, &people1)

// 	fmt.Println("---- Print Field ----")
// 	fmt.Println(people1.Name)
// 	fmt.Println(people1.Height)
// 	fmt.Println(people1.Mass)
// 	fmt.Println(people1.HairColor)
// 	fmt.Println(people1.Films[0])
// }

// MEMBACA RESPONSE NESTED OBJECT JSON
func main() {
	response, err := http.Get("https://swapi.dev/api/people")
	if err != nil {
		log.Fatal("error consume API")
	}

	responseData, errBody := io.ReadAll(response.Body)
	if errBody != nil {
		log.Fatal("error read response body")
	}
	defer response.Body.Close()

	var people1 PeopleResponse
	json.Unmarshal(responseData, &people1)

	fmt.Println("---- Print Field ----")
	fmt.Println(people1.Count)
	for _, value := range people1.Results {
		fmt.Println("nama:", value.Name)
		fmt.Println("hair:", value.HairColor)
		// fmt.Println("hair:", value.Films[0])
		if len(value.Films) > 0 {
			for _, v := range value.Films {
				fmt.Println("films:", v)

			}
		}
		fmt.Println("------------------")
	}
}
