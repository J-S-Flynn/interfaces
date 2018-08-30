package main

import (
	"fmt"
	"net/http"
	"os"
)

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

	url := "http://google.com"

	getConection(url)

}

//this func utalises the interface type bot. so once again we are infering the struct type
func printGreating(b bot) {

	fmt.Println(b.getGreating())
}

//not a great deal of custom logic in the two get greting funcs but there is enough for us to understand that
//some functions need to be defined in a custom way. its also important to note that the recivers have no variable only a type
func (englishBot) getGreating() string {

	return "Hello there you beauties"
}

func (spanishBot) getGreating() string {

	return "Ola allí bellezas"
}

func getConection(url string) {

	resp, error := http.Get(url)
	if error != nil {
		fmt.Println("no responce from the url")
		os.Exit(1)
	}

	fmt.Println(resp.Status)

}
