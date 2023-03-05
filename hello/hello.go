package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func testQuote() {
	fmt.Println(quote.Go())
}

func testHello() {
	msg, _ := greetings.Hello("golang")
	fmt.Println(msg)
}

func testHellos() {
	names := []string{"foo", "bar", "baz"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

func testHelloHanleError() {
	msg2, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg2)
}

func main() {
	testHello()
	testHellos()
}
