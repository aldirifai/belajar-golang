package main

import "fmt"

func main() {
	var nilai = 90
	var absensi = 80

	var lulusNilai bool = nilai > 80
	var lulusAbsensi bool = absensi > 80

	var lulus = lulusNilai && lulusAbsensi
	fmt.Println(lulus)

	fmt.Println(nilai > 80 && absensi > 80)
}
