package main

import "fmt"


// separate domain code from outside world, side effect

const spanish = "spanish"
const engc = "Hello, "
const spanishc = "Hola, "

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishc + name
	}
	return engc + name
}

func main() {
	fmt.Println(hello("Sai", "spanish"))
}