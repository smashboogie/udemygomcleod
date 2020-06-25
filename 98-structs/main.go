//aggragate data type - compose together values of different type

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   26,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.last, p1.first, p1.age)
	fmt.Println(p2.last, p2.first, p2.age)
}
