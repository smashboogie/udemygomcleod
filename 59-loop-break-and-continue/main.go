package main

//keyword - break
//keyword - continue
import (
	"fmt"
)

func main() {
	// operator '/' = simple division, rounded
	//x := 83 / 40
	// operator '%' = remainder (called modulo)
	//	y := 83 % 40
	//fmt.Println(x, y)

	z := 1
	for {
		z++
		if z > 100 {
			break
		}
		if z%2 != 0 {
			continue
		}
		fmt.Println(z)

	}
}
