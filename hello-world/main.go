package main

import "fmt"

const (
	spanish = "ESP"
    french = "FRA"
    deutsch = "GER"

    spanishHelloPrefix = "Hola, "
    englishHelloPrefix = "Hello, "
    frenchHelloPrefix = "Bonjour, "
    deutschHelloPrefix = "Hallo, "
)

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
	fmt.Println(Hello("Rodrigo", "ESP"))
}