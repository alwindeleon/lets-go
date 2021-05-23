package main

import "fmt"

const filipino = "filipino"
const spanish = "spanish"
const englishHelloPrefix = "Hello, "
const filipinoHelloPrefix = "Oi, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case filipino:
		prefix = filipinoHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Bai", "spanish"))
}
