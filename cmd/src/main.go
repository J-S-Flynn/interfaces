package main

import "fmt"

type bot interface {
	getGreating() string
}

//the structs below have no unique fields to them but creating type englishBot and spanBot as a type struct allow
//for the use or functions with them
type englishBot struct{}

type spanishBot struct{}

func main() {

	//interfaces have a fw functions, but one of the things interfaces can help us with is code eusibility.

	eb := englishBot{}
	sb := spanishBot{}

	printGreating(eb)
	printGreating(sb)
}

func printGreating(b bot) {

	fmt.Println(b.getGreating())
}

//not a great deal of custom logic in the two get greting funcs but there is enough for us to understand that
//some functions need to be defined in a custom way
func (eb englishBot) getGreating() string {

	return "Hello there you beauties"
}

func (sp spanishBot) getGreating() string {

	return "Ola allí bellezas"
}
