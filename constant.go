package main

import "fmt"

func main() {
	const (
		firstName string = "Fadjrir"
		lastName         = "Herlambang"
		value            = 1000
	)

    // Uncomment this line bellow, to see error.
    // firstName = "Ananda"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}