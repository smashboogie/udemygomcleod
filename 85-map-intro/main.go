//map is an unordered list
// maps associate values of one type (key) with values of
//another type (element or values)
//maps hold references to an underlying data structure
//if you pass a map to a function that changes the contents of a map, the
//changes will be visible to the caller
//maps can be constructied using the usual literal syntax
//to distimnguish a missing entry- you use the "comma ok" idiom

package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      32,
		"Miss Money": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])
	//enter a key that has no value stored in the map (below) it returns 0 value
	fmt.Println(m["Barnabus"])
	// you can check if a value exists in the map too (below) using ', ok '
	v, ok := m["Barnabus"]
	fmt.Println(v)
	fmt.Println(ok)
	// can be added into an IF statment (common in go) below
	if v, ok := m["James"]; ok {
		fmt.Println("This is the IF", v)
	}
}
