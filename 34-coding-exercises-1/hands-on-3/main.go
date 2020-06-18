package main

import (
	"fmt"
)

const (
	a = 42
)

func main() {
	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%#X\n", a)
	b := (a << 1)
	fmt.Printf("%d\n", b)
	fmt.Printf("%b\n", b)
	fmt.Printf("%#X\n", b)

}
