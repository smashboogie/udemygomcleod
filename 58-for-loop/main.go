package main

//keyword = if

import "fmt"

func main() {
	x := 1
	// for statment with a single condition
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

	//if statment - if some condition is true/false, we do something with it

	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++

	}
}
