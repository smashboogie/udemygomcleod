package main

import "fmt"

func main() {
	//anonymous - no name
	s1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Steph",
		last:  "Thorpe",
		age:   54,
	}
	s2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Val",
		last:  "Fredericks",
		age:   58,
	}

	fmt.Println(s1)
	fmt.Println(s2)
}
