package main

import (
	"fmt"
)

func main() {
	a := 42
	b := 43
	fmt.Println(a == b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)

	//OR
	c := (42 == 42)
	fmt.Println(c)
}
