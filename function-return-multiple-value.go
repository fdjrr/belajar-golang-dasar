package main

import "fmt"

func getFullName() (string, string) {
	return "Fadjrir", "Herlambang"
}

func main() {
	firstName, _ := getFullName()
    
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)
}