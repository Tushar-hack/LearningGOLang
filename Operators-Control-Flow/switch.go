package main

import "fmt"

func main() {
	i := 10
	// fallthrough Keyword : to force the execution flow through successive case block.
	switch i {
	case -5:
		fmt.Println("-5")
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("20")
		fallthrough
	default:
		fmt.Println("Default")
	}

	// switch with conditions

	a := 10
	b := 20

	switch {
	case a+b == 30:
		fmt.Println("Sum is equal to 30.")
	case a+b <= 30:
		fmt.Println("Sum is less than or equal to 30.")
	default:
		fmt.Println("Sum is greater than 30.")
	}
}
