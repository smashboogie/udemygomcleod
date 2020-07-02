//functions about about being modular (small chunks)
//syntax = func (r reciever) identif ier (parameters) (returns(s)) {...}
//know the difference between parameters and arguments
//we define our func with parameters (if any)
//we call our func and pass in arguments (if any)
//everything in GO is PASS BY VALUE (what you see is what you get)

package main

import (
	"fmt"
)

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

//basic function
func foo() {
	fmt.Println("hello from foo")
}

//passing an argument into a function
func bar(s string) {
	fmt.Println("Hello", s)

}

//return values into a function
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

//multiple return values into a function
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}
