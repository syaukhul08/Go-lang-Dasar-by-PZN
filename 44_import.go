package main

import (
	"belajar_golang_dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Syaukhul")
	// helper.sayGoodBye("Syaukhul") --- Error
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) --- Error
}
