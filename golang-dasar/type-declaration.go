package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noNIK NoKTP = "12345678"
	var marriedStatus Married = false

	fmt.Println(noNIK)
	fmt.Println(marriedStatus)
}
