package main

import "fmt"

func modify(a int) {
	a += 100
}

func main() {
	// Function is called directly by passing the value of the variable as an argument
	// the parameter is copied into another location of your memory
	// so when accessing or modifying the variable within your function , only the copy is accessed or modified
	// the original value never modified.
	a := 10
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
}
