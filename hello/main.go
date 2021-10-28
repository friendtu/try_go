package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings log: ")
	log.SetFlags(0)
	messages, err := greetings.Hellos([]string{"Tud", "Lian"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
