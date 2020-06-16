package main

//control flow
//sequence (top to bottom)
//iterative (loop)
//conditional (one of two directions based on a condition)

import (
	"fmt"
)

func main() {
	// true/false = pre-declared constant
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	//adding expressions below
	if 2 == 2 {
		fmt.Println("005")
	}
	//initionalization statement inside an IF statment =  (;)
	//x variable below is limited in scope (to IF) to avoid clashing elesewhere
	if x := 42; x == 42 {
		fmt.Println("006")
	}
}
