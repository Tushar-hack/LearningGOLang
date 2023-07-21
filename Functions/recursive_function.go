package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	/*
		 	--> Recursion is a concept where a function calls itself by direct or indirect means
			--> The function keeps on calling itself until it reaches a base case
			--> Used to solve a problem where the solution is dependent on the smaller instance of the same problem
	*/
	n := 5
	result := factorial(n)
	fmt.Println("Factorial of", n, "is :", result)
}
