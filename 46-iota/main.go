package main

//iota is a predeclared identifier

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
	d
)

//starts over with new const set (0, etc)
const (
	e = iota
	f = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n%T\n%T\n%T\n%T\n%T\n", a, b, c, d, e, f)

}
