//slicing a slice is a colon :
package main

import (
	"fmt"
)

func main() {
	//Composite Literal below
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	//slices value of a single indext in the array
	fmt.Println(x[1:])
	//slices values of a range of indexes (up to but not including 3)
	fmt.Println(x[1:3])
	for i, v := range x {
		fmt.Println(i, v)
	}
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
