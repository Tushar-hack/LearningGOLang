package main

import (
	"fmt"
	"strconv"
)

func main() {
	// First way of TypeCasting
	var i int = 90
	fmt.Printf("Type of i is %T. \n", i)

	var f float64 = float64(i)
	fmt.Printf("Type of f is %T. \n", f)

	// Second way of TypeCasting (string -> int or vice versa)
	var s string = strconv.Itoa(i)
	fmt.Printf("Type of s is %T. \n", s)

	x, y := strconv.Atoi(s) // Atoi return two things first the value and second the error if any.
	fmt.Printf("Value is %v and Error is %v", x, y)
}
