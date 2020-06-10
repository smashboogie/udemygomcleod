/*
keyword - switch
https://golang.org/ref/spec#Switch_statements

"Switch" statements provide multi-way execution. An expression or type specifier is compared to the "cases" inside the "switch" to determine which branch to execute.
There are two forms: expression switches and type switches. In an expression switch, the cases contain expressions that are compared against the value of the switch expression

In an expression switch, the switch expression is evaluated and the case expressions, which need not be constants, are evaluated left-to-right and top-to-bottom
 the first one that equals the switch expression triggers execution of the statements of the associated case; the other cases are skipped.

If no case matches and there is a "default" case, its statements are executed. There can be at most one default case and it may appear anywhere in the "switch" statement.

 A missing switch expression is equivalent to the boolean value true.
 example below
*/
package main

import (
	"fmt"
)

func main() {

	switch {
	case false:
		fmt.Println("false should not print")
	case true:
		fmt.Println("true is default - should print")
	}
}
