package main

import (
	"fmt"
)

func random() interface{} {
	return 100
}

func main() {
	// result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}

}
