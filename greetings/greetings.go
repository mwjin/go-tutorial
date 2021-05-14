package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("The name is empty")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hello, %v!",
		"Welcome, %v!",
		"Hi, %v!",
		"Bye, %v!",
	}

	return formats[rand.Intn(len(formats))]
}
