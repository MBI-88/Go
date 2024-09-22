package main

/*
A var declaration creates a variable of a particular type, attaches a name to it,
and sets its initial value. Each declaration has the general form
var name type = expression

Either the type or the = expression part may be omitted, but not both. If the type is omitted,
it is determined by the initializer expression. If the expression is omitted, the initial value is
the zero value for the type, which is 0 for numbers, false for booleans, "" for str ings, and nil
for interfaces and reference types (slice, pointer, map, channel, function). The zero value of an
aggregate type like an array or a struct has the zero value of all of its elements or fields


Within a function, an alternate form called a short variable declaration may be used to declare
and initialize local variables. It takes the form name := expression, and the type of name is
determined by the type of expression. Here are three of the many short variable declarations

Ke ep in mind that := is a declarat ion, where as = is an assig nment. A multi-variable 
declaration should not be confused with a tuple assignment (§2.4.1), in which each variable on the
lef t-hand side is assigned the corresponding value from the right-hand side:
i, j = j, i // swap values of i and j

One subtle but important point: a short variable declaration does not necessarily declare all the
var iables on its left-hand side. If some of them were already declared in the same lexical block
(§2.7), then the short variable declaration acts like an assignment to those variables.

A short variable declaration must declare at least one new variable, however, so this code will
not compile:
f, err := os.Open(infile)
// ...
f, err := os.Create(outfile) // compile error: no new variables

A short variable declaration acts like an assignment only to variables that were already
declared in the same lexical block; declarations in an outer block are ignored.
*/

import "fmt"


// Declaracion de tipo y de valor (declaracion a nivel de paquetes)
var i, k, j int
var b, f, s = true, 2.3, "four"

func variableDeclaration(){
	time := "8:30 am"
	valor1 := 20
	valor2 := 30
	result := (valor1 + valor2)

	fmt.Printf("Resultado de variables locales => %s -- %d -- %d -- %d\n",time,valor1,valor2,result)


}


func main() {
	variableDeclaration()

}
