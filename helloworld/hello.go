package main

import "fmt"

func getGreetingPrefix(language string) (prefix string) {
	const englishHelloPrefix = "Hello"
	const spanishHelloPrefix = "Hola"
	const frenchHelloPrefix = "Bonjour"

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := getGreetingPrefix(language)

	return prefix + " " + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
