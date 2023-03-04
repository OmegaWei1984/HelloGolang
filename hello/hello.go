package main

import (
    "fmt"
    "log"
    "rsc.io/quote"
    "example.com/greetings"
)

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
    msg, _ := greetings.Hello("golang")
    fmt.Println(msg)
    msg2, err := greetings.Hello("")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(msg2)
}

