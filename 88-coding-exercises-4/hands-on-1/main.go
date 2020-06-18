//Using a COMPOSITE LITERAL:
//create an ARRAY which holds 5 VALUES of TYPE int
//assign VALUES to each index position
//Range over the array and print the values out.
//Using format printing - print out the TYPE of the array
//difference between array and slice format
// array [5]int{}
// slice []int{}

package main

import "fmt"

func main() {
	//Composite Literal array below and assigning values
	x := [5]int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}

//printing out the type of the array with Printf
