package main

import "fmt"

func main() {
	// An array is a collection of similar data elements stored at contiguous memory locations.
	/*
	 -> Fixed Length
	 -> Element should be of same data Type.
	*/

	// Syntax -> var <array_name> [<size>] <data_type>

	//var grades [5]int = [5] int{10,20,30}  --> this is long way
	//grades := [5] int{10,20,30} --> short way
	grades := [...]int{10, 20, 30, 40, 50} // ellipses --> length == no of values in the array
	fmt.Println(grades)
	fmt.Printf("Length of an array is %v\n", len(grades))

	// Change or assign a value at index
	fmt.Println("Indexes in array")
	fmt.Println(grades[2])
	grades[2] = 45
	fmt.Println(grades[2])

	// looping through array
	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}
	// using range keyword
	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}

	// Multi-dimensional Array ( 2D - Array )
	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr[2][1])
}
