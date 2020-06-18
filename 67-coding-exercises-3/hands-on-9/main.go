//Create a program that uses a switch statement with the switch expression
//specified as a variable of TYPE string with the IDENTIFIER “favSport”.
//hint - IDENTIFIER is before (above) switch statment
//hint - cases must match type (e.g strings, ints, etc)
package main

import (
	"fmt"
)

func main() {
	favSport := "soccer"
	switch favSport {
	case "football":
		fmt.Println("its football!")
	case "baseball":
		fmt.Println("its baseball!")
	case "soccer":
		fmt.Println("its soccer")
	}
}
