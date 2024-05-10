package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Fadjrir"
	sayHelloTo(firstName, "Herlambang")
	sayHelloTo("Ananda", "Maharani")
}