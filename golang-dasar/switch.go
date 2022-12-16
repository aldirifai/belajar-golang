package main

import "fmt"

func main() {
	name := "Aldi"

	switch name {
	case "Aldi":
		fmt.Println("Hello Aldi")
	case "Rifai":
		fmt.Println("Hello Rifai")
	default:
		fmt.Println("Hi, Kenalang dong !")
	}

	//switch lenght := len(name); lenght > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	lenght := len(name)

	switch {
	case lenght < 2:
		fmt.Println("Nama terlalu sedikit")
	case lenght > 5:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
