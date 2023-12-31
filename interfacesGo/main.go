package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct {
}
type spanishBot struct {
}
type binaryBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	return "Hello there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (binaryBot) getGreeting() int {
	return 11
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
