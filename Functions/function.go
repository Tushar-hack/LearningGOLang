package main

import (
	"fmt"
)

func addNumbers(a int, b int) int { //a and b are params
	sum := a + b
	return sum
}
func operation(a int, b int) (int, int) {
	sum := a + b
	difference := a - b
	return sum, difference
}

// return named values
//
//	func operation(a int, b int) (sum int,difference int) {
//		sum = a + b
//		difference = a - b
//		return
//	}

func main() {
	// Self-contained units of code which carry out a certain job
	// Help us divide a program into small manageable, repeatable and organisable chunks.
	// Reusability and Abstraction

	//Declaration := func <function_name>(<params>) <return type> { body of the function }

	// Calling a function
	sum := addNumbers(10, 20) // 10 and 20 are arguments
	fmt.Println(sum)

	// Return Types := single value , multiple value, named value
	sum, difference := operation(20, 10)
	fmt.Println(sum, difference)

}
