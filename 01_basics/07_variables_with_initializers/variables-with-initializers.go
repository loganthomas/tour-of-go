/*
Variables with initializers
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
*/
package main

import (
	"fmt"
	"reflect"
)

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	// fmt.Println(i, j, c, python, java)
	fmt.Printf("Type of i     : %T Value of i: %v\n", i, i)
	fmt.Printf("Type of j     : %T Value of j: %v\n", j, j)
	fmt.Printf("Type of c     : %T Value of c: %v\n", c, c)
	fmt.Printf("Type of python: %T Value of python: %v\n", python, python)
	fmt.Printf("Type of java  : %T Value of java: %v\n", java, java)
	fmt.Println()
	fmt.Printf("Type of i     : %v Value of i: %v\n", reflect.TypeOf(i), i)
	fmt.Printf("Type of j     : %v Value of j: %v\n", reflect.TypeOf(j), j)
	fmt.Printf("Type of c     : %v Value of c: %v\n", reflect.TypeOf(c), c)
	fmt.Printf("Type of python: %v Value of python: %v\n", reflect.TypeOf(python), python)
	fmt.Printf("Type of java  : %v Value of java: %v\n", reflect.TypeOf(java), java)
}
