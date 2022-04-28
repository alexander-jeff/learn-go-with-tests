package main

import (
	"fmt"
)

const (
	englishHelloPrefix string = "Hello, "
	spanishHolaPrefix  string = "Hola, "
	SPANISH            string = "Spanish"
)

func Hello(name string, language string) string {
	return getPrefix(language) + getName(name)
}

func getPrefix(language string) string {
	switch language {
	case SPANISH:
		return spanishHolaPrefix
	default:
		return englishHelloPrefix
	}
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
