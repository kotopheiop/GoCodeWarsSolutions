package main

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}

func main() {
	fmt.Println(Greet("Gopher")) // Hello, Gopher how are you doing today?
}
