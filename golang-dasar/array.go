package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhamad"
	names[1] = "Aldi"
	names[2] = "Rifai"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		10, 20, 30,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var data [10]string
	fmt.Println(len(data))
}
