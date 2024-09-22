package main

/*
	Here are four ways to declare a function with two parameters and one result, all of type int.
	The blank identifier can be used to emphasize that a parameter is unused.

	func add(x int, y int) int { return x + y }
	func sub(x, y int) (z int) { z = x - y; return }
	func first(x int, _ int) int { return x }
	func zero(int, int) int { return 0 }
	fmt.Printf("%T\n", add) // "func(int, int) int"
	fmt.Printf("%T\n", sub) // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero) // "func(int, int) int"

	The type of a function is sometimes called its signature. Two functions have the same type or
	signature if they have the same sequence of parameter types and the same sequence of result
	types. The names of parameters and results don’t affect the type, nor does whether or not they
	were declared using the factored form.

	Every function call must provide an argument for each parameter, in the order in which the
	parameters were declared. Go has no concept of default parameter values, nor any way to
	specify arguments by name, so the names of parameters and results don’t matter to the caller
	except as documentation.
	Parameters are local variables within the body of the function, with their initial values set to
	the arguments supplied by the caller. Function parameters and named results are variables in
	the same lexical block as the function’s outermost local variables.

	
*/

import (
	"math"
	"fmt"
)

func hypot(x,y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main(){
	fmt.Println(hypot(3,4))


}