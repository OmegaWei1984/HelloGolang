package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func hello() {
	msg, _ := greetings.Hello("golang")
	fmt.Println(msg)
}

func hellos() {
	names := []string{"foo", "bar", "baz"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

func helloHanleError() {
	msg2, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg2)
}

func main() {
	hello()
	hellos()
}
