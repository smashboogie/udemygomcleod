//create your raw string literals first (xs1, xs2)
//Create a slice of a slice of string ([][]string)
// Store the following data in the multi-dimensional slice:
//"James", "Bond", "Shaken, not stirred"
//"Miss", "Moneypenny", "Helloooooo, James."
//Range over the records,
//range over the data in each record.

package main

import (
	"fmt"
)

func main() {
	//raw string literals first - xs1: = []string{"string1", "string2"}
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	//print them
	fmt.Println(xs1)
	fmt.Println(xs2)
	//slice of a slice of a string AND ([][]string{value1, value2})
	//stored multidimensional slice
	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)
	//range OVER with for loop to give index positions
	//FOR #1 - create a for loop (i = index and range )
	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}

//FOR #2 - create for loop to give values (j = index and range)
