package main

import "fmt"

func main() {
	var name1 = "Fadjrir Herlambang"
	var name2 = "Ananda Maharani"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}