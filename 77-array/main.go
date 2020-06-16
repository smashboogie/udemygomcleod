//array is a building block in go func()
//zero(0) based index
//primarily used for building blocks of slices
//use slices instead
//array is numbered sequence of elements of a single type
//lenght is a part of the array type


package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}
