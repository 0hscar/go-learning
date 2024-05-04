package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Pär", "Lars", "Tjong"}

	messages, err := greetings.Hellos(names) //Does not return error, if "" returns error

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
