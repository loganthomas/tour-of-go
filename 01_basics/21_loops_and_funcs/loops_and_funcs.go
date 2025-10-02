/*
Exercise: Loops and Functions
As a way to play with functions and loops, let's implement a square root function:
given a number x, we want to find the number z for which z² is most nearly x.

Computers typically compute the square root of x using a loop.
Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

z -= (z*z - x) / (2*z)

Repeating this adjustment makes the guess better and better
until we reach an answer that is as close to the actual square root as can be.

Implement this in the func Sqrt provided.
A decent starting guess for z is 1, no matter what the input.

To begin with, repeat the calculation 10 times and print each z along the way.
See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.

Hint: To declare and initialize a floating point value, give it floating point syntax or use a conversion:

z := 1.0
z := float64(1)

Next, change the loop condition to stop once the value has stopped changing
(or only changes by a very small amount).
See if that's more or fewer than 10 iterations.
Try other initial guesses for z, like x, or x/2.
How close are your function's results to the math.Sqrt in the standard library?

(Note: If you are interested in the details of the algorithm,
the z² − x above is how far away z² is from where it needs to be (x),
and the division by 2z is the derivative of z²,
to scale how much we adjust z by how quickly z² is changing.

This general approach is called Newton's method.
It works well for many functions but especially well for square root.)
*/

package main

import (
	"fmt"
	"math"
)

// attempt 1 (for loop with incremental index)
//
//	func Sqrt(x float64) float64 {
//		z := 1.0
//		for i := 0; i < 10; i++ {
//			fmt.Println("start z", z)
//			z = z - ((z*z - x) / (2 * z))
//			fmt.Println("next z", z)
//		}
//		return z
//	}
//
// attempt 2 (for loop with explicit slice of numbers)
//
//	func Sqrt(x float64) float64 {
//		z := 1.0
//		iterations := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//		// this will error since not using i
//		// for i := range iterations {
//		// this will error since declaring no new variables on left side of := (use = instead)
//		// for _ := range iterations {
//		for _ = range iterations {
//			fmt.Println("start z", z)
//			z = z - ((z*z - x) / (2 * z))
//			fmt.Println("next z", z)
//		}
//		return z
//	}
//
// attempt 3 (for loop with a slice of length 10 with indices 0-9 )
//
//	func Sqrt(x float64) float64 {
//		z := 1.0
//		for _ = range make([]int, 10) {
//			// for _ = range [10]int{} {
//			fmt.Println("start z", z)
//			z = z - ((z*z - x) / (2 * z))
//			fmt.Println("next z", z)
//		}
//		return z
//	}
//
// final (range over in go 1.22+; my version is go1.24.2 darwin/arm64)
func Sqrt(x float64) float64 {
	z := 1.0
	for range 10 {
		z -= ((z*z - x) / (2 * z))
	}
	return z
}

func earlySqrt(x float64) float64 {
	z := 1.0
	iter := 0
	const eps = 1e-10
	diff := 1.0

	for diff > eps {
		newZ := z - ((z*z - x) / (2 * z))
		diff = math.Abs(z - newZ)
		z = newZ
		iter++
	}
	fmt.Printf("Converged after %d iterations\n", iter)
	return z
}

// robust earlySqrt
func robustEarlySqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("cannot calculate square root of negative number: %v", x)
	}

	z := 1.0
	iter := 0
	const (
		epsilon = 1e-10
		maxIter = 100
	)
	diff := 1.0

	for diff > epsilon && iter < maxIter {
		newZ := z - (z*z-x)/(2*z)
		diff = math.Abs(z - newZ)
		z = newZ
		iter++
	}

	if iter == maxIter {
		return z, fmt.Errorf("failed to converge after %d iterations", maxIter)
	}

	return z, nil
}

func main() {
	fmt.Println(
		Sqrt(25), earlySqrt(25), math.Sqrt(25),
		Sqrt(2), earlySqrt(2), math.Sqrt(2),
	)
	result, err := robustEarlySqrt(-2)
	fmt.Println(result, err)
	// notice the = and not :=
	result, err = robustEarlySqrt(2)
	fmt.Println(result, err)
}
