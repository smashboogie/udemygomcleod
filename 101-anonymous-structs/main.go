//aggragate data type - compose together values of different type

package main

import "fmt"

func main() {
	//anonymous - no name
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
}
