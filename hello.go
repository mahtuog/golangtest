package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	languagePrefix := greetingPrefix(language)

	return languagePrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
