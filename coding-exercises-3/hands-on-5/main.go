//Print out the remainder (modulus) which is found for
//each number between 10 and 100 when it is divided by 4.
// used a for loop
//modulo = i%4
package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("I divide %v by 4 and get the remainder which is %v\n", i, i%4)
	}
}
