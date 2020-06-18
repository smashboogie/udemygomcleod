//Create a program that uses a switch statement with no switch expression specified.
//default switch with no expression (blank) is true.

package main

import (
	"fmt"
)

func main() {
	switch {
	case true:
		fmt.Println("its twue its twue!")
	case false:
		fmt.Println("its false")
	}
}
