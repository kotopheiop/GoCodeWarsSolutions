package main

import "fmt"

func Greet(name string) string {
	if name == "Johnny" {
		return fmt.Sprint("Hello, my love!")
	}

	return fmt.Sprintf("Hello, %v!", name)
}

func main() {
	fmt.Println(Greet("Alfred")) // Hello, Alfred!
	fmt.Println(Greet("Johnny")) // Hello, my love!

}
