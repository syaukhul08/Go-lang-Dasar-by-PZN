package main

import (
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHai() {
	fmt.Println("Hello, My Name is", customer.Name, "I live in", customer.Address, "and I'm ", customer.Age, "year old")
}

func main() {
	var budi Customer
	budi.Name = "Budi Yono"
	budi.Address = "Indonesia"
	budi.Age = 25

	budi.sayHai()
}
