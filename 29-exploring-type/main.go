package main

import (
	"fmt"
)

var y = 42

//DECLARED the VARIABLE with the IDENTIFER 'z'
// is of TYPE string
// And I am ASSIGN the value 'Shaken, not Stirred'

var z string = "Shaken, not Stirred"
var a string = `James, said, "Shaken not stirred"'`

// GO is STATIC - not a DYNAMIC language
// a VARIABLE is DECLARED to hold a VALUE of a certain type

func main() {
	fmt.Println(y)
	fmt.Printf("%T \n", y)
	fmt.Println(z)
	fmt.Printf("%T \n", z)
	fmt.Println(a)
	fmt.Printf("%T \n", a)
}
