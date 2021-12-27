package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	ebot := englishBot{}
	sbot := spanishBot{}
	printGreeting(ebot)
	printGreeting(sbot)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// these functions are only for example purpose. We assume that there is custom logic involved in both greeting functions

func (englishBot) getGreeting() string {
	return "Hey there!"
}

func (spanishBot) getGreeting() string {
	return "Holla!"
}
