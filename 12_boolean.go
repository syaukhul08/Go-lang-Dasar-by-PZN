package main

import "fmt"

func main() {
	var ujian = 88
	var absensi = 88

	var lulusUjian = ujian > 80
	var lulusAbsensi = absensi > 88

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

}
