package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Glady")
	fmt.Println(message)
}
