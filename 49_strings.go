package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Syaukhul A'la", "Syaukhul"))
	fmt.Println(strings.Contains("Syaukhul A'la", "Heru"))

	fmt.Println(strings.Split("Syaukhul A'la", " "))

	fmt.Println(strings.ToLower("Syaukhul A'la"))

	fmt.Println(strings.ToUpper("Syaukhul A'la"))

	fmt.Println(strings.Trim("      Syaukhul A'la          ", " "))

	fmt.Println(strings.ReplaceAll("Joko Heru Joko ", "Joko", "Andi"))
}
