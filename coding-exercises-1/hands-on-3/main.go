package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

// default values are called ZERO VALUES
func main() {

	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)

}
