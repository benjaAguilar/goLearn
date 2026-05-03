package main

import "fmt"

const greeting = "Hello "

func SayHello(name string) string {

	if name == "" {
		name = "World!"
	}

	return greeting + name
}

func main() {
	fmt.Println(SayHello(""))
}
