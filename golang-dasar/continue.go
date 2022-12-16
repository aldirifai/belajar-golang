package main

import "fmt"

func main() {

	// bilangan ganjil
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke ", i)
	}

	// bilangan genap
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		fmt.Println("Perulangan ke ", i)
	}
}
