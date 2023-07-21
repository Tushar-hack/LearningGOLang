package main

import "fmt"

func main() {
	/*
		--> An anonymous function is a function that is declared without any identifier to refer it
		-->	 They can accept inputs and return outputs, just as standard functions do
		--> They can be  used for containing functionality that need not to be named and possibly for a short-term use.
	*/
	x := func(l int, b int) int {
		return l * b
	}
	fmt.Printf("%T \n", x)
	fmt.Println(x(20, 30))

	y := func(l int, b int) int {
		return l * b
	}(20, 60)
	fmt.Println(y)
	fmt.Printf("%T \n", y)
}
