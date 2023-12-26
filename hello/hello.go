package main

import "fmt"
import "example.com/greetings"
import "log"

func main ()  {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Ankit")
	if err!= nil {
        log.Fatal(err)
    }

	fmt.Println(message)
	
}
