package main

import "fmt"

func main() {
	//First way of Declaring Variable
	// var <var_name> <dataType> = <value>
	var name string = "Tushar"
	var surname string = "Tak"
	var rollNumber int = 223
	var student bool = true
	var percentage float64 = 82.8
	fmt.Println(name, surname, rollNumber, student, percentage)

	//shorthand way ! Work with variables of same data type
	var my, mum string = "Rikku", "Maya"
	fmt.Println(my)
	fmt.Println(mum)

	//For different types variables
	var (
		nickname string = "Rahul"
		age      int    = 21
	)

	fmt.Println(nickname, age)

	//Short Variable Declaration
	brother := "Bhavesh"
	// brother = 12 -> this will give you error cause it only store String now
	fmt.Print(brother)
}
