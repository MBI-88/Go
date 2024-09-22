package main

/*
	Pointers
	A pointer value is the address of a variable. A pointer is thus the location at which a value is
	stored. Not every value has an address, but every variable does. With a pointer, we can read
	or update the value of a variable indirectly, without using or even knowing the name of the
	variable, if indeed it has a name.

	If a variable is declared var x int, the expression &x (‘‘address of x’’) yields a pointer to an
	integer variable, that is, a value of type *int, which is pronounced ‘‘pointer to int.’’ If this
	value is called p, we say ‘‘p points to x,’’ or equivalently ‘‘p contains the address of x.’’ The variable to which p points is written *p. 
	The expression *p yields the value of that variable, an
	int, but since *p denotes a variable, it may also appear on the left-hand side of an assignment,
	in which case the assignment updates the variable

	The zero value for a pointer of any type is nil. The test p != nil is true if p points to a variable. 
	Pointers are comparable; two pointers are equal if and only if they point to the same
	variable or both are nil.

	Each time wetake the address of a variable or copy a pointer, we create new aliases or ways to
	identify the same variable. For example, *p is an alias for v. Pointer aliasing is useful because
	it allows us to access a variable without using its name, but this is a double-edged sword: to
	find all the statements that access a variable, we have to know all its aliases. It’s not just pointers 
	that create aliases; aliasing also occurs when we copy values of other reference types like
	slices, maps, and channels, and even structs, arrays, and interfaces that contain these types.

	Pointers are key to the flag package, which uses a program’s command-line arguments to set
	the values of certain variables distributed throughout the program. To illustrate, this variation
	on the earlier echo command takes two optional flags: -n causes echo to omit the trailing
	newline that would normally be printed, and -s sep causes it to separate the output arguments by the contents of 
	the string sep instead of the default single space.
*/

import (
	"fmt"
	"flag"
	"strings"
)


func Pointers(){
	x := 1 
	p := &x // p, puntero de tipo entero (apunta a x)
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	// Igualdad de punteros
	var a, b int
	fmt.Println(&a == &a, &a == &b, &a == nil)
	
}

func ReturnPointer() *int {
	/* Cada llamada a la funcion retorna un direccion distinta */
	variable := 1
	return &variable
}

func incr(p *int) int {
	*p++
	return *p
}

func flagParse() {
	n := flag.Bool("n",false,"omit tralling newline")
	sep := flag.String("s"," ","separator")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n {
		fmt.Println()
	}
}

func main(){
	//Pointers()
	//p := ReturnPointer()
	//fmt.Println(p)

	//v := 1
	//incr(&v) // el valor de v es 2
	//fmt.Println(incr(&v)) // el valor de v es 3

	flagParse()

}