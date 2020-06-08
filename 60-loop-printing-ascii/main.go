package main

import (
	"fmt"
)

func main() {
	// for ALWAYS "init; condition; post {}"
	for i := 33; i <= 122; i++ {

		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}
