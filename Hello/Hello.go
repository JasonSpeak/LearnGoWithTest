package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const EnglishHello = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return GetLanguage(language) + name
}

func GetLanguage(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = EnglishHello

	}
	return
}

func main() {
	fmt.Println(Hello("Tom", spanish))
}
