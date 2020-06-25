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
type secretAgent struct {
	person
	ltk bool
}

func main() {
	//outer type is secretAgent
	sa := secretAgent{
		//inner type person gets promoted to outer type
		person: person{
			first: "James",
			last:  "Bonds",
			age:   32,
		},

		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   26,
	}

	fmt.Println(sa)
	fmt.Println(p2)
	fmt.Println(sa.first, sa.last, sa.age, sa.ltk)
	fmt.Println(p2.last, p2.first, p2.age)
}
