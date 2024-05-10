package main

import "fmt"

func sumAll(firstName string, numbers ...int) int {
    fmt.Println(firstName)

	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll("Fadjrir", 10, 90, 30, 50, 40)
	fmt.Println(total)

	slice := []int{10,20,30,40,50}
	total = sumAll("Ananda", slice...)
	fmt.Println(total)
}