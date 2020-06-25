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
			"chocolate",
			"martini",
			"rum and coke"},
	}
	p2 := person{
		first: "Miss",
		last:  "Monnypenny",
		favFlavors: []string{
			"vanmilla",
			"strawberry",
			"pumpkin"},
	}
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("________")
	}
}
