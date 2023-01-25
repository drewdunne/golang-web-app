package main //This is the name of *this* package (not an import!)

import "fmt"

func main() {

	var i int
	var whatToSay string = "Hello World"

	i = 7

	fmt.Println(whatToSay, i)

	whatWasSaid, whatWasAlsoSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, whatWasAlsoSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
