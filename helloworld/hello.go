package main

import "fmt"

func Hello(name string, language string) string {

	const greetingEnglish = "Hello, "
	const greetingFrench = "Bonjour, "
	const greetingSpanish = "Hola, "

	if name == "" {
		name = "world"
	}

	switch language {
	case "spanish":
		return greetingSpanish + name
	case "french":
		return greetingFrench + name
	case "english":
		return greetingEnglish + name
	default:
		return greetingEnglish + name
	}
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
