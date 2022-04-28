package main

import (
	"fmt"
)

const (
	englishGreeting string = "Hello, "
	spanishGreeting string = "Hola, "
	frenchGreeting  string = "Bonjour, "

	SPANISH string = "Spanish"
	FRENCH  string = "French"
)

func Hello(name string, language string) string {
	return getPrefix(language) + getName(name)
}

func getPrefix(language string) (prefix string) {
	switch language {
	case SPANISH:
		prefix = spanishGreeting
	case FRENCH:
		prefix = frenchGreeting
	default:
		prefix = englishGreeting
	}
	return
}

func getName(name string) string {
	if name == "" {
		return "World!"
	}
	return name
}

func main() {
	fmt.Println(Hello("", ""))
}
