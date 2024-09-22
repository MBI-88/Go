package main 

/*
	Another way to create a variable is to use the built-in function new. The expression new(T)
	creates an unnamed variable of type T, initializes it to the zero value of T, and returns its
	address, which is a value of typ e *T

	p := new(int) // p, of type *int, points to anunnamed int variable
	fmt.Println(*p) // "0"
	*p = 2 // sets the unnamed int to 2
	fmt.Println(*p) // "2"

	There is one exception to this rule: two variables whose type carries no information and is
	therefore of size zero, such as struct{} or [0]int, may, depending on the implementation,
	have the same address.

	Since new is a predeclared function, not a keyword, it’s possible to redefine the name for
	something else within a function, for example:

	func delta(old, new int) int { return new - old }
*/

import "fmt"

func newInt() *int {
	return new(int)
}

func newIntVar() *int {
	var dummy int
	return &dummy
}

func useNew(){
	// Cada llamada a new retorna una nueva variable con una unica dirección
	p := new(int)
	q := new(int)
	fmt.Println(p == q) // "false"
}


func main(){
	// Estas dos funciones hacen lo mismo que new
	fmt.Printf("newInt => %d newIntVar => %d\n",*newInt(),*newIntVar())
	useNew()

}