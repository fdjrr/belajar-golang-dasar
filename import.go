package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Fadjrir")
	// helper.sayGoodbye("Fadjrir") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
