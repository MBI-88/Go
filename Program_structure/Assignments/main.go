package main

/*
	The value held by a variable is updated by an assignment statement, which in its simplest form
	has a variable on the left of the = sign and an expression on the right.
	x = 1 // named variable
	*p = true // indirect variable
	person.name = "bob" // struct field
	count[x] = count[x] * scale // array or slice or map element
	count[x] *= scale

	Numeric variables can also be incremented and decremented by ++ and -- statements:
	v := 1
	v++ // same as v = v + 1; v becomes 2
	v-- // same as v = v - 1; v becomes 1 again

	Tuple Assignment

	Anot her form of assignment, known as tuple assignment, allows several variables to be
	assigned at once. All of the right-hand side expressions are evaluated before any of the variables are updated, 
	making this form most useful when some of the variables appear on both
	sides of the assignment, as happens, for example, when swapping the values of two var iables:
	x, y = y, x
    a[i], a[j] = a[j], a[i]

	Assignability

	As signment statements are an explicit form of assignment, but there are many places in a
	program where an assignment occurs implicitly: a function call implicitly assigns the argument
	values to the corresponding parameter variables; a return statement implicitly assigns the
	return operands to the corresponding result variables; and a literal expression for a composite
	type (§4.2) such as this slice:

	medals := []string{"gold", "silver", "bronze"}

	Wh ether two values may be compared with == and != is related to assignabilit y: in any comparison, 
	the first operand must be assignable to the type of the second operand, or vice versa.
	As with assignability.
*/

import "fmt"


func asignTuple(x,y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x,y = y,  x + y
	}
	return x
}

func main(){
	// Asignación con tupla
	fmt.Printf("Resultado %d\n",asignTuple(10,5))
	fmt.Printf("Fib => %d\n",fib(5))

	

}