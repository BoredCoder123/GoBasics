package greetings

import "fmt"
import "errors"

func Hello(name string) (string, error) {

	if name == ""{
		return "", errors.New("Name cannot be empty")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
