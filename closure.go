package main

import "fmt"

func main() {
	name := "Fadjrir"
	counter := 0

	increment := func() {
		name := "Ananda"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}