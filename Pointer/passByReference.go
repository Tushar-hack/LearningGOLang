package main

import "fmt"

func modifyMaps(m map[string]int) {
	m["K"] = 75
}

func modifySlice(s []int) {
	s[0] = 90
}

func modifying(a *int) {
	*a += 56
}

func main() {
	// address of the variable is passed into the function call as the actual parameter
	// all the operations are performed on the value stored at the address of the actual parameter

	a := 24
	fmt.Println(a)
	modifying(&a)
	fmt.Println(a)

	//slices and maps are passed by reference , by default

	slice := []int{10, 20, 30}
	fmt.Println(slice)
	modifySlice(slice)
	fmt.Println(slice)

	ascii_codes := make(map[string]int)
	ascii_codes["A"] = 65
	ascii_codes["F"] = 70
	fmt.Println(ascii_codes)
	modifyMaps(ascii_codes)
	fmt.Println(ascii_codes)
}
