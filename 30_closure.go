package main

import "fmt"

func main() {
	name := "Syaukhul"
	increment := func() {
		name = "A'la"
		fmt.Println("increment")
	}

	increment()

	fmt.Println(name)
}
