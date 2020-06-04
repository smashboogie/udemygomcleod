package main

import (
	"fmt"
)

// int for whole numbers without decimals is 30
var a int

// float64 for decimal numbers ie 30.3
var b float64

//int8 signed example-negative ie. -128 to 127
var c int8 = -128

//uint8 unsigned example-zero ie 0 to 255 (alias=byte)
//int32 (alias=rune)

func main() {
	a = 7
	b = 2.30000
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	//short declaration operator
	x := 42
	y := 42.34334
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

}
