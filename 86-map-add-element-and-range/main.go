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

	//adding an element to the map
	m["todd"] = 33

	// can be added into an IF statment (common in go) below
	if v, ok := m["James"]; ok {
		fmt.Println("This is the IF", v)
	}
	//getting the range of the map
	for k, v := range m {
		fmt.Println(k, v)
	}
	//with a slice
	xi := []int{4, 5, 7, 8, 9, 42}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
