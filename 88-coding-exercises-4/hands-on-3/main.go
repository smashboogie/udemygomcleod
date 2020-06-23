//Using a COMPOSITE LITERAL:
//create a SLICE of TYPE int
//assign 10 VALUES
//hint the only difference between an array and slice is the
// array [5]int{}
// slice []int{}

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}
