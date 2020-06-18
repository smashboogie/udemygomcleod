package main

import (
	"fmt"
)

//untyped constant
const a = 42

//typed constant
const b int = 43

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)
}
