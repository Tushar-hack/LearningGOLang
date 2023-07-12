package main

import "fmt"

func main() {
	//printing String
	fmt.Print("Hello World !")

	//printing variable
	var name string = "Tushar"
	fmt.Print(name)

	//printing variable and String
	var surname string = "Tak"
	fmt.Print("This is", name, surname)

	//using println to print in a new line
	fmt.Println(name)
	fmt.Println(surname)

	//using the printf ( Format Specifier )
	// 1. %v -> formats the value in default format
	fmt.Printf("My name is %v", name)
	// 2. %d -> formats decimal integers
	var grades int = 42
	fmt.Printf("Marks: %d", grades)
}

// Some Specific format Specifier
/*
%T -> Type of value
%c -> character
%q -> quoted characters/string
%s -> plain String
%t -> true or false
%f -> floating numbers
%.2f -> floating number upto 2 decimal
*/
