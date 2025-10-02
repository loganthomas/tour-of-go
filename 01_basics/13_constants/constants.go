/*
Constants

Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.
*/

package main

import "fmt"

const Pi = 3.14

var varPi float64 = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// this will cause an error since Pi is defined as a constant
	// Pi = 3.1

	// this is fine since varPi is defined as a variable
	varPi = 3.1
	fmt.Println("varPi", varPi)
}
