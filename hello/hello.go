package main

import (
    "fmt"
    "log"
    "rsc.io/quote"
    "example.com/greetings"
)


func testQuote() {
    fmt.Println(quote.Go())
}

func testHello() {
    msg, _ := greetings.Hello("golang")
    fmt.Println(msg)
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
}

