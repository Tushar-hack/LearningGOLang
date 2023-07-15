package main

import "fmt"

func main() {
	// Slice: Continuous segment of an underlying array.
	// variable Typed (elements can be added or removed)
	// more flexible
	// Slice 3 component : Pointer, length -> len(), capacity -> cap()

	//declaration := <slice_name> = [] <data_type>{<values>}

	slice := []int{10, 20, 30}

	fmt.Println(slice)
	// creating a slice from an array := array[ start_index : end_index ]
	//end_index is not included
	// start_index is 0 by default AND end_index is size of an array

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_1 := arr[1:8]
	fmt.Println(slice_1)

	// capacity difference
	fmt.Println(cap(arr))
	fmt.Println(cap(slice_1))

	sub_slice := slice_1[1:4]
	fmt.Println(sub_slice)

	// Second way of making slice := make([]<data_type>, length, capacity)

	slice_2 := make([]int, 5, 8)
	fmt.Println(slice_2)
	fmt.Println(len(slice_2))
	fmt.Println(cap(slice_2))

	// Changes in slice will affect the array
	slice_1[0] = 110
	fmt.Println(arr)
	fmt.Println(slice_1)

	// Appending to a slice := func append
	slice_1 = append(slice_1, 120, 130, 450, 67)
	fmt.Println(slice_1)

	// If the capacity of the slice is exceeding the original capacity then it will create a new array with the capacity 2X of original_capacity
	//and return that array
	fmt.Println(cap(slice_1))

	// Appending slice to a slice
	new_slice := append(sub_slice, slice_1...)
	fmt.Println(new_slice)

	// deleting from a slice
	new_Arr := [5]int{10, 20, 30, 40, 50}
	i := 2
	fmt.Println(new_Arr)
	firstSlice := new_Arr[:i]
	secondSlice := new_Arr[:i+1]
	third_slice := append(firstSlice, secondSlice...)
	fmt.Println(third_slice)

	//Copying from a slice
	src_slice := []int{10, 20, 30, 40, 50}
	dest_slice := make([]int, 3)
	num := copy(dest_slice, src_slice)
	fmt.Println(dest_slice)
	fmt.Println("Numbers of element copied: ", num)

	// looping through a slice
	for index, value := range src_slice {
		fmt.Println(index, "=> ", value)
	}

	for _, value := range src_slice {
		fmt.Println(value)
	}
}
