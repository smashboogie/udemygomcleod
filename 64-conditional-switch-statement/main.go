package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should also not print")
	case (3 == 3):
		fmt.Println("this should print")
		//go has no default fallthrough; it stops at first case
		//must be specifically defined (as below)
		fallthrough
	case (4 == 4):
		fmt.Println("this is true but it only prints with fallthrough")
		fallthrough
	case (7 == 9):
		fmt.Println("not true1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
		//default case is when nothing else evauluates to true
		fallthrough
	default:
		fmt.Println("this is default case")
	}
	// switch using a string
	switch "Bond" {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("Bond is the case that hit")
	case "Q":
		fmt.Println("This is not bond - its Q")
	default:
		fmt.Println("default case for Bond String")
	}
	// switch using assigning a value (n) to the string
	//multiple values in a switch
	n := "Oddjob"
	switch n {
	case "Oddjob", "Dr. No", "Tattoo":
		fmt.Println("Bond Enemies")
	case "M":
		fmt.Println("M")
	case "Q":
		fmt.Println("Q")
	default:
		fmt.Println("default case the enemy value")
	}
}
