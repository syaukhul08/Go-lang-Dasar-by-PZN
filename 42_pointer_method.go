package main

import (
	"fmt"
)

type Man struct {
	Name string
}

//tambah pointer supaya data bisa berubah
func (man *Man) Merried() {
	man.Name = "Mr. " + man.Name
}

func main() {

	syaukhul := Man{"Syaukhul"}
	syaukhul.Merried()

	fmt.Println(syaukhul.Name)
}
