package main

import (
    "fmt"

    "github.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Giancarlo")
    fmt.Println(message)
}
