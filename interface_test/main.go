package main

import "fmt"

type bot interface {
	getGreeting() string
}

// empty struct types declared
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// this func accept an argument of type bot
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println((b.getGreeting()))
}

// englishBot implements the bot interface prototype and is thus a bot
func (englishBot) getGreeting() string {
	// very custom logic for generating engish greeting
	return "Hello"
}

// spanishBot implements the bot interface prototype and is thus a bot
func (spanishBot) getGreeting() string {
	// very custom logic for generating spanish greeting
	return "Hola"
}
