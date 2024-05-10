package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Fadjrir Herlambang", "Fadjrir"))
	fmt.Println(strings.Contains("Fadjrir Herlambang", "Budi"))

	fmt.Println(strings.Split("Fadjrir Herlambang", " "))

	fmt.Println(strings.ToLower("Fadjrir Herlambang"))
	fmt.Println(strings.ToUpper("Fadjrir Herlambang"))
	fmt.Println(strings.ToTitle("Fadjrir Herlambang"))

	fmt.Println(strings.Trim("      Fadjrir Herlambang     ", " "))
	fmt.Println(strings.ReplaceAll("Fadjrir Joko Fadjrir", "Fadjrir", "Budi"))
}
