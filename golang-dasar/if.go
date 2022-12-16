package main

import "fmt"

func main() {
	var name = "Aldi"

	if name == "Aldi" {
		fmt.Println("Hello Aldi")
	} else if name == "Rifai" {
		fmt.Println("Hello Rifai")
	} else {
		fmt.Println("Hi, Kenalan dong !")
	}

	// long version
	//var lenght = len(name)
	//
	//if lenght > 5 {
	//	fmt.Println("Nama terlalu panjang")
	//} else {
	//	fmt.Println("Nama sudah benar")
	//}

	// sort version
	if lenght := len(name); lenght > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
