package main

import (
	"fmt"
	"io"
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

	//we are going to make a byte slice in a new way here, not the way we have seen before i.e bs := []byte{}
	//make takes a type , in this case a byte and initalises it with 99999 elements inside it, so we could
	//make a byte slice as aboove wich will grow and shrinc, but this will make an empty byte slice and make it n big
	// bs := make([]byte, 99999)

	//now we can get the responce body and use the read reciver on that passing in the byte slice that we created.
	//the read function will then push the responce into the byte slice and we can then print out that bs by rappint it in a string
	//conversion.
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	fmt.Println("\n\n new print out")

	//however go has a lot of helper methods so we could simply write
	io.Copy(os.Stdout, resp.Body)

	//its importent to undertand that this is useing the writer interface rather than the read.

}
