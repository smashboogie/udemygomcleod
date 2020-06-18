package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	//apending to a slice
	x = append(x, 77, 78, 110)
	fmt.Println(x)

	y := []int{234, 434, 340}
	x = append(x, y...)
	fmt.Println(x)
	//deleting from a slice
	//full index = [4 5 7 8 42 77 78 110 234 434 340]
	//this should keep the first two #s (4, 5), remove 2 (7,8) and keep rest (42+)
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
