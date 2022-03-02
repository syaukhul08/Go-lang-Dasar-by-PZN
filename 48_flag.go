package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "put your host")
	var user *string = flag.String("user", "root", "put your user")
	var password *string = flag.String("password", "root", "put your password")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
}
