//Create a for loop using this syntax = for {}
//Have it print out the years you have been alive
//hint - it uses break
//hint - it uses if

package main

import (
	"fmt"
)

func main() {
	bd := 1972
	for {
		if bd > 2020 {
			break
		}
		fmt.Println(bd)
		bd++

	}
}
