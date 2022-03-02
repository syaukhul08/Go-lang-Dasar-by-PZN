package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	utc := time.Date(2022, 3, 2, 2, 10, 20, 20, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02" //atau time.RFC3339
	parse, _ := time.Parse(layout, "2001-03-08")
	fmt.Println(parse)
}
