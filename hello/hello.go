package main

import (
	"fmt"
	"log"

	"github.com/mwjjeong/go-tutorial/greetings"
)

func main() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
