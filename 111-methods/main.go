//A method is nothing more than a FUNC attached to a TYPE.

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
}
