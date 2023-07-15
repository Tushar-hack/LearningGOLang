package main

import (
	"fmt"
)

func main() {
	// Maps := Unordered collection of key/value Pairs
	// Implemented by Hash Tables
	// efficient add, get and delete operations

	// Declaring := var <map_name> map[<key_data_type>]<value_data_type>  { <key-value-pairs> }
	// Zero value of map is nil.

	myMap := map[int]string{
		1: "Tushar",
		2: "Shubhang",
	}
	fmt.Println(myMap)

	// Accessing a map
	fmt.Println(myMap[2])

	// Gives length of the map
	fmt.Println(len(myMap))

	// Getting a key
	value1, found1 := myMap[3]
	fmt.Println(value1)
	fmt.Println(found1)

	// making maps using make(map[key]value, <initial_capacity>) function
	//second argument is optional

	myMap2 := make(map[string]string)
	fmt.Println(myMap2)
	value, found := myMap2["we"]
	fmt.Println(value)
	fmt.Println(found)

	// Adding, Updating key-value pair
	myMap2["En"] = "English"
	myMap2["Hi"] = "Hindi"
	myMap2["fr"] = "French"
	fmt.Println(myMap2)

	// Deleting a key-value pair
	delete(myMap2, "En")
	fmt.Println(myMap2)

	// Iterating over map
	for key, value := range myMap2 {
		fmt.Println(key, "=>", value)
	}

	// truncating a map means removing all elements
	// first method is by iterating and deleting all key-value one by one

	//for key ,_ := range myMap2 {
	//	delete(myMap2, key)
	//}

	// second method is initializing it with a new empty map
	//myMap2 = make(map[string]string)
}
