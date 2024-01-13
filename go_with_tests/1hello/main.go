package main

import "fmt"


// separate domain code from outside world, side effect

const (
	spanish = "spanish"
	french = "french"
	
	engc = "Hello, "
	spanishc = "Hola, "
	frenchc = "Bonjour, "
)

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return addPrefix(language) + name
}

func addPrefix(language string) (prefix string){
	switch language{
	case french:
		prefix = frenchc
	case spanish:
		prefix = spanishc
	default:
		prefix = engc
	}

	return 
}

func main() {
	fmt.Println(hello("Sai", "spanish"))
}