package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("j([a-z])o")

	fmt.Println(regex.MatchString("jiko"))
	fmt.Println(regex.MatchString("jono"))
	fmt.Println(regex.MatchString("jacKo"))

	search := regex.FindAllString("aan iin uun een", 2)
	fmt.Println(search)

}
