package main

import "fmt"


// separate domain code from outside world, side effect

const c = "Hello, "
func hello(name string) string{
	if name == "" {
		name = "World"
	}
	return c + name
}

func main() {
	fmt.Println(hello("Sai"))
}