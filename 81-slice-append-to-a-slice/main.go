//what append does is append the elements to the end of the slice
//and return the result
//
//func append (slice []T, elements ...T) []T
//... = unlimited VALUES of that type (variadic)
//T is a placeholder for any given type

package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 78, 110)
	fmt.Println(x)

	y := []int{234, 434, 340}
	x = append(x, y...)
	fmt.Println(x)
}
