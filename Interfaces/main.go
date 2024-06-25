package main

import "fmt"

type bot interface {
	getGreeting() string
}

type spanishBot struct{}

type englishBot struct{}

func main() {
	fmt.Println("Main")
	eb := englishBot{}
	sb := spanishBot{}
	// Since the struct has got the method getGreeting they are honorary bot
	// so they can be used
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (s spanishBot) getGreeting() string {
	return "Holaa"
}

func (e englishBot) getGreeting() string {
	return "Hello World"
}
