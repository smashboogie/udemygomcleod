package main

import (
	"fmt"
)

//unfurling a slice of int
func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("The total is", x)

}
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	return sum
}

//you can also pass in zero or more values
//https://play.golang.org/p/Qc5sq_AK_T
//	x := sum()
//	fmt.Println("The total is", x)

//func main() {
//	x := sum("james")
//	fmt.Println("The total is", x)
//}

//variadic parameter has to be the final parameter
//https://play.golang.org/p/euQ8PDQ8RN
//below - string is first - variadic parameter HAS TO BE LAST
//func sum(s string, x ...int) int {
