package main

import (
	"fmt"
	"log"

	"github.com/mwjjeong/go-tutorial/greetings"
)

func main() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Minwoo Jeong")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"L style", "M style", "N style"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
