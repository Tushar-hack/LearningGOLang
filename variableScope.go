package main

import "fmt"

var name string = "Tushar Tak" //Global variable
func main() {
	//variable scope is defined using blocks
	// outer block
	city := "Ajmer"
	fmt.Println(name)
	{
		// inner block
		country := "India"
		fmt.Println(city)
		fmt.Println(country)
		fmt.Println(name)
	}
	fmt.Println(city)
	// fmt.Println(country) this line will give error because no such variable exist in outer scope
}
