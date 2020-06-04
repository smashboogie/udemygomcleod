package main

import (
	"fmt"
)

func main() {
	//start with decimal 4
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	//bit shifting below from 4 to
	//moving decimal place over 1 (from 100 to 1000)
	y := x << 1
	fmt.Printf("%d\t\t%b\n\n", y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	//binary digits shift 10x each statment
	fmt.Printf("%d\t%b\n", kb, kb)
	fmt.Printf("%d\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}
