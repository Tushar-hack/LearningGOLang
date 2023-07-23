package main

import "fmt"

func main() {
	// Pointer is a variable that holds memory address of another variable.
	// they point where the memory is allocated and provide ways to find or even change the value located at the memory address.

	// Declaring Pointer := var <pointer_name> *<data_type>
	var num int = 10
	// Initializing a pointer
	var p *int = &num

	/*
		--> & operator := address of a variable can be obtained by preceding the name of a variable with an ampersand sign (&)
		--> * operator := dereference operator when placed before address return the value at that address
	*/
	fmt.Println(p)
	fmt.Println(*p)
	*p = 86
	fmt.Println(*p)
}
