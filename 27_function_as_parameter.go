package main

import "fmt"

//Using function type declarations
type Filter func(string) string

//function as parameter change to Filter
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", filter(nameFiltered))
}

//function as parameter
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello ", filter(nameFiltered))
// }

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Syaukhul", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

}
