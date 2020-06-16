//composite data type -
// in go its called a struct

package main

import (
	"fmt"
)

func main() {
	//Composite Literal below
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
  //accessing values by index position below
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
  //go over range of array and get values with a for loop below
	for i, v := range x {
		fmt.Println(i, v)
	}

}

//SLICE allows you to groupo together values of the same type.
