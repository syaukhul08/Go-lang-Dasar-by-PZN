package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Semarang", "Jawa Tengah", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Jepara"

	//tanpa operator *
	//address2 = &Address{"Sleman", "DIY", "Indonesia"}

	//dengan operator
	*address2 = Address{"Sleman", "DIY", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	var address3 *Address = new(Address)
	address3.City = "Jakarta"
	fmt.Println(address3)
}
