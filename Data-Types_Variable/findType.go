package main

import (
	"fmt"
	"reflect"
)

func main() {
	// %T format specifier
	i := 60
	fmt.Printf("Variable i is of type %T. \n", i)

	//using TypeOf function
	s := "Tushar"
	fmt.Printf("Type of variable s is %v.", reflect.TypeOf(s))
}
