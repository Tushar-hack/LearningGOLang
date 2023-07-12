package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i * i)
	}

	// some advance looping strategy

	j := 1
	for j <= 5 {
		fmt.Println(j * j)
		j += 1
	}

	// Infinite Loop
	// for { logic here } -> This loop will run forever.

	// break := ends the loop immediately when encountered
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	//continue := skips the current iteration of loop
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
