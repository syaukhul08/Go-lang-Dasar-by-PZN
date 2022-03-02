package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Eko", 30},
		{"Joko", 40},
		{"Heru", 24},
	}

	//sebelum sort
	fmt.Println("sebelum : ", users)
	sort.Sort(UserSlice(users))
	//setelah sort
	fmt.Println("setelah : ", users)
}
