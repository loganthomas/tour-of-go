/*
Short variable declarations
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/
package main

import (
	"fmt"
	"reflect"
)

var logan int = 5
var hudson = 42

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Printf("Type of i: %v (value: %v)\n", reflect.TypeOf(i), i)
	fmt.Printf("Type of j: %v (value: %v)\n", reflect.TypeOf(j), j)
	fmt.Printf("Type of k: %v (value: %v)\n", reflect.TypeOf(k), k)
	fmt.Printf("Type of c: %v (value: %v)\n", reflect.TypeOf(c), c)
	fmt.Printf("Type of python: %v (value: %v)\n", reflect.TypeOf(python), python)
	fmt.Printf("Type of java: %v (value: %v)\n", reflect.TypeOf(java), java)
}
