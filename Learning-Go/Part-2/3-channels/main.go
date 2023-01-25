package main

import (
	"encoding/json"
	"fmt"
	"log"

	helpers "github.com/drewdunne/golang-intro"
)

const numPool = 10

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
} 

func main() {
	// imagine we get this JSON from an external API
	myJson := `
	[ 
		{
			"first_name" : "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name" : "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]
`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled);

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("Error Marshalling Json", err)
	}

	log.Println("Marshalled:")
	fmt.Println(string(newJson))
}

func CalculateValue(intChan chan int) { // channel of int is the type
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}


