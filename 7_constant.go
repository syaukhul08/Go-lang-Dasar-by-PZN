package main

import "fmt"

func main() {

	// const fristName string = "Syaukhul"
	// const lastName = "A'la"
	// const value = 1000

	//Multiple Constant
	const (
		fristName string = "Syaukhul"
		lastName         = "A'la"
		value            = 1000
	)

	fmt.Println(fristName)
	fmt.Println(lastName)
	fmt.Println(value)

	//error
	// fristName = "Tidak Bisa diubah"
	// lastName = "Tidak bisa diubah"
}
