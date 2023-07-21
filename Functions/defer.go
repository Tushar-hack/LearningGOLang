package main

import "fmt"

func printName(str string) {
	fmt.Println(str)
}
func printRollNo(str int) {
	fmt.Println(str)
}
func printAddr(str string) {
	fmt.Println(str)
}
func main() {
	/*
		--> A defer statement delays the execution of function until the surrounding function returns
		--> The deferred call's arguments are evaluated immediately , but the function call is not executed until the surrounding function returns.
	*/
	printName("Tushar")
	defer printRollNo(100)
	printAddr("Jaipur Road")
}
