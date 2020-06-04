package main

import (
	"fmt"
)

//const is a keyword
//const untyped is without type below
// const noted in separate expressions
const a = 42
const b = 47.8
const c = "James Bond"

//const typed is with type below
// all const in a single expression
const (
	d int     = 43
	e float64 = 43.8
	f string  = "Peter Sellers"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n%T\n%T\n", a, b, c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n%T\n%T\n", d, e, f)

}
