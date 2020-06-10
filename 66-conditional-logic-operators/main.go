/*
conditional logic operators
fmt.Println(true && true) = both conditions are true
fmt.Println(true && false) = one condition can be true AND one can be false
fmt.Println(true || true) = either/OR one condition can be true
fmt.Println(true || false) = one could be true OR one could be false
fmt.Println(!true)

EXAMPLE below
fmt.Println(true && true)
fmt.Println(true && false)
fmt.Println(true || true)
fmt.Println(true || false)
fmt.Println(!true)
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("true && true\t %v\n", true && true)
	fmt.Printf("true && false\t %v\n", true && false)
	fmt.Printf("true || true\t %v\n", true || true)
	fmt.Printf("true || false\t %v\n", true || false)
	fmt.Printf("!true\t\t %v\n", !true)
	fmt.Printf("!false\t\t %v\n", !false)

}
