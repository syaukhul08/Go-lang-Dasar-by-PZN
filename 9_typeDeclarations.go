package main

import "fmt"

func main() {
	type NoKTP string
	type Merried bool

	var myNoKTP NoKTP = "13193192344"
	var merriedStatus Merried = false
	fmt.Println(myNoKTP)
	fmt.Println(merriedStatus)
}
