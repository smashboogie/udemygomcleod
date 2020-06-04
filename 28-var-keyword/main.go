package main

import (
	"fmt"
)

//vars can be used outside of a main function
// DECLARE the variable 'y'
// ASSIGN the variable 45
// DECLARE and ASSIGN = INTITILIZATION
var y = 45

//DECLARE there is a VARIABLE with the IDENTIFIER 'z' and is of TYPE IDENTIFIER
// ASSIGNS the ZERO VALUE of TYPE into to 'z'
//false for booleans, 0 for numeric types, "" for strings,
//and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	//Short declaration operator can only be used inside a function
	// DECLARE a variable and ASSIGN a VALUE at the same time
	x := 42
	fmt.Println(x)

	fmt.Println(y)
}
