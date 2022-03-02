package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Syaukhul"
	names[2] = "A'la"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		98,
		34,
		56,
	}
	fmt.Println(values)

	//function array
	fmt.Println(len(names))
	fmt.Println(len(values))

}
