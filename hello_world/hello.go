package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const german = "German"

const englishPrefixHello = "Hello, "
const spanishPrefixHello = "Hola, "
const frenchPrefixHello = "Bonjour, "
const germanPrefixHello = "Hallo, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishPrefixHello
	case french:
		prefix = frenchPrefixHello
	case german:
		prefix = germanPrefixHello
	default:
		prefix = englishPrefixHello
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
