package main

import (
	"log"
	"github.com/drewdunne/golang-intro"
)

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepered",
	}

	PrintInfo(&dog)
	var myVar helpers.SomeType = helpers.SomeType{TypeName: "hello", TypeNumber: 1}
	log.Println(myVar)
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}
