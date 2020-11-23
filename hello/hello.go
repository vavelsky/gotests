package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const italian = "Italian"
const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "
const italianHelloPrefix = "Ciao "

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	default:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Pawel", ""))
}
