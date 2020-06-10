package main

import (
	"fmt"
)

func main() {
	x := 46
	if x == 40 {
		fmt.Println("our value was 40")

	} else if x == 41 {
		fmt.Println("our value was 41")
	} else if x == 43 {
		fmt.Println("our value was 43")
	} else if x == 44 {
		fmt.Println("our value was 44")
	} else if x == 45 {
		fmt.Println("our value was 45")
	} else {
		fmt.Println("our value was not 40, 41, 43, 44, or 45")
	}
}
