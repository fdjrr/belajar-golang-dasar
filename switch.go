package main

import "fmt"

func main() {
	name := "Fadjrir"

	switch name {
	case "Fadjrir":
		fmt.Println("Hello Fadjrir")
	case "Ananda":
		fmt.Println("Hello Ananda")
	default:
		fmt.Println("Kenalan Donk")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}