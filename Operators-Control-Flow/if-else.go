package main

import "fmt"

func main() {
	// if example
	var a = "happy"
	if a == "happy" {
		fmt.Println(a)
	}
	// if-else example
	fruit := "Grapes"
	if fruit == "Apple" {
		fmt.Println("Fruit is Apple.")
	} else {
		fmt.Println("Fruit is not Apple.")
	}

	// Nested if-else example
	name := "Tushar"
	if name == "Bhavesh" {
		fmt.Println("This is bhavesh Here.")
	} else if name == "Tushar" {
		fmt.Println("This is Tushar Here.")
	} else {
		fmt.Println("Unknown Person.")
	}
}
