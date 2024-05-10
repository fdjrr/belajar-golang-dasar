package main

import "fmt"

func getFullName() (firstName, lastName string) {
	firstName = "Fadjrir"
	lastName = "Herlambang"

	return
}

func main() {
	a, b := getFullName()
	fmt.Println(a)
	fmt.Println(b)
}