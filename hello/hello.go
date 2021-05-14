package main

import (
	"fmt"

	"github.com/mwjjeong/go-tutorial/greetings"
)

func main() {
	message := greetings.Hello("Minwoo Jeong")
	fmt.Println(message)
}
