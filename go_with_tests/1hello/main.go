package main

import "fmt"


// separate domain code from outside world, side effect


func hello() string{
	return "Hello, World!"
}

func main() {
	fmt.Println(hello())
}