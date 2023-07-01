package main

import "fmt"

func main() {
	//Zero values are those default values which store in a variable when we don't store anything.
	var s int
	var f float64
	var ns string
	var tf bool
	fmt.Printf("%d\n", s)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%s\n", ns)
	fmt.Printf("%t\n", tf)
}
