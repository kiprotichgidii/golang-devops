package main

import "fmt"

const spanish = "ESP"
const french = "FRA"
const deutsch = "GER"
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const deutschHelloPrefix = "Hallo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case deutsch:
		prefix = deutschHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Gedion", "GER"))
}
