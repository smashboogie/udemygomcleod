//Using a COMPOSITE LITERAL:
//create a SLICE of TYPE int
//assign 10 VALUES
//hint the only difference between an array and slice is the
// array [5]int{}
// slice []int{}

package main

import (
	"fmt"
)

func main() {
	x := []int{000, 100, 200, 300, 400, 500, 600, 700, 800, 900}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}
