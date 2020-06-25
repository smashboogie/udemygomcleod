//aggragate data type - compose together values of different type

package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"pump", "capp", "spice",
		},
	}

	p2 := person{
		first:      "Miss",
		last:       "Moneypenny",
		favFlavors: []string{"mango", "ube", "vanilla"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
		fmt.Println(p2.first)
		fmt.Println(p2.last)
		for i, v := range p2.favFlavors {
			fmt.Println(i, v)
		}
	}
}
