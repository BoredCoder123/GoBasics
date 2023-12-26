package greetings

import "fmt"
import "errors"
import "math/rand"

func Hello(name string) (string, error) {

	if name == ""{
		return "", errors.New("Name cannot be empty")
	}

	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// return message, nil

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
    formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v!, well met!",
	}

	return formats[rand.Intn(len(formats))]
}
