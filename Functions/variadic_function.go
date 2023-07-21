package main

import (
	"fmt"
)

func printDetails(student string, subject ...string) {
	fmt.Println("hey ", student, ", here are your subjects - ")
	for _, sub := range subject {
		fmt.Println(sub)
	}
}
func sumNumbers(numbers ...int) (sum int) {
	sum = 0
	for _, value := range numbers {
		sum += value
	}
	return
}
func main() {
	//variadic function := functions that accepts variable number of arguments
	// It is possible to pass a varying number of arguments of the same type as referenced in the function signature.
	// declaration := func addNumber(str String,numbers ...int) int { }
	fmt.Println(sumNumbers(10, 20, 30, 40, 50))
	printDetails("Joe", "Physics", "Biology")

	// Blank identifier ( _ )
}
