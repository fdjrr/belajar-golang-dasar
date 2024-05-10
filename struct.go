package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu(){
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var cust Customer
	cust.Name = "Fadjrir"
	cust.Address = "Indonesia"
	cust.Age = 18

	cust.sayHi("Ananda")
	cust.sayHuuu()

	//fmt.Println(cust.Name)
	//fmt.Println(cust.Address)
	//fmt.Println(cust.Age)
	//
	//joko := Customer{
	//	Name:    "Joko",
	//	Address: "Cirebon",
	//	Age:     35,
	//}
	//fmt.Println(joko)
	//
	//budi := Customer{"Budi", "Jakarta", 35}
	//fmt.Println(budi)
}