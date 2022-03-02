package main

import "fmt"

func endApp() {

	//recover
	messege := recover()
	if messege != nil {
		fmt.Println("Error dengan messege : ", messege)
	}

	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Syaukhul")
}
