package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtp NoKTP = "1271020512040001"
	var marriedStatus Married = false
	fmt.Println(noKtp)
	fmt.Println(marriedStatus)
}