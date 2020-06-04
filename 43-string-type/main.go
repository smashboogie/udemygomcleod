package main

import (
	"fmt"
)

// all code in GO is UTF-8
// GO strings are immutable (can't change)
// You CAN assign a new value
func main() {
	//double quotes doesn't include carriage returns etc
	s := "James Bond"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	//raw string literal - backticks `` include everything including returns
	t := `James
  Bond`
	fmt.Println(t)
	fmt.Printf("%T\n", t)
	//can convert string to slice of bytes below
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
}
