package main

import "fmt"

func main() {
	var name = "Fadjrir"

	if name == "Fadjrir" {
		fmt.Println("Hello Fadjrir")
	} else if name == "Ananda" {
		fmt.Println("Hello Ananda")
	} else if name == "Ery" {
		fmt.Println("Hello Ery")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}