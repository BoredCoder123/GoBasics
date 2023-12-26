package main

import "fmt"
import "example.com/greetings"
import "log"

func main ()  {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, err := greetings.Hello("Ankit")
	// if err!= nil {
    //     log.Fatal(err)
    // }

	// fmt.Println(message)

	names := []string{"Ankit", "Kattu", "Kathait"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for key, val := range messages {
		fmt.Println(key + ": " + val)
	}
}
