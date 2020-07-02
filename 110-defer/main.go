package main

import (
	"fmt"
)

//unfurling a slice of int
func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
