package hello

const (
	englishGreeting  string = "Hello, "
	spanishGreeting  string = "Hola, "
	frenchGreeting   string = "Bonjour, "
	hawaiianGreeting string = "Aloha, "

	SPANISH  string = "Spanish"
	FRENCH   string = "French"
	HAWAIIAN string = "Hawaiian"
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
	case HAWAIIAN:
		prefix = hawaiianGreeting
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
