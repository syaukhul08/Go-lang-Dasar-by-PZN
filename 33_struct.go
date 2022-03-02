package main

import (
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	var syaukhul Customer
	syaukhul.Name = "Syaukhul"
	syaukhul.Address = "Indonesia"
	syaukhul.Age = 20

	fmt.Println(syaukhul.Name)
	fmt.Println(syaukhul.Address)
	fmt.Println(syaukhul.Age)

	budi := Customer{
		Name:    "Budi",
		Address: "Semarang",
		Age:     43,
	}
	fmt.Println(budi)
}
