package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	msg := greetings.HelloGreeting("Jael")
	fmt.Println(msg)
}