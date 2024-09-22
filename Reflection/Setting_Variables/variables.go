package main

import (
	"fmt"
	"reflect"
)

/*
	Recall that some Go expressions like x, x.f[1], and *p denote variables, but others like x+1
	and f(2) do not. A variable is an addressable storage location that contains a value, and its
	value may be updated through that address.

	The value within a is not addressable. It is merelyacopy of the integer 2. The same is true of
	b. The value within c is also non-addressable, being a copy of the pointer value &x. In fact, no
	reflect.Value returned by reflect.ValueOf(x) is addressable. But d, derived from c by
	dereferencing the pointer within it, refers to a variable and is thus addressable. We can use
	this approach, calling reflect.ValueOf(&x).Elem(), to obtain an addressable Value for any
	variable x.

	We obtain an addressable reflect.Value whenever we indirect through a pointer, even if we
	started from a non-addressable Value. All the usual rules for addressability have analogs for
	reflection. For example, since the slice indexing expression e[i] implicitly follows a point er, it
	is addressable even if the expression e is not. By analogy, reflect.ValueOf(e).Index(i)
	refers to a variable, and is thus addressable even if reflect.ValueOf(e) is not.

	To recover the variable from an addressable reflect.Value requires three steps. First, we call
	Addr(), which returns a Value holding a pointer to the variable. Next, we call Interface()
	on this Value, which returns an interface{} value containing the pointer. Finally, if we
	know the type of the variable, we can use a type assertion to retrieve the contents of the interface
	as an ordinary pointer. We can then update the var iable through the pointer:


*/

func main(){
	x := 2        // value type variable?
	a := reflect.ValueOf(2) // 2 int no
	b := reflect.ValueOf(x) // 2 int no
	c := reflect.ValueOf(&x) // &x *int no
	d := c.Elem() // 2 int yes (x)
	fmt.Printf("Dado %d\n a=%v\n b=%v\n c=%v\n d=%v\n",x,a,b,c,d)

	y := 5
	z := reflect.ValueOf(y).Elem() //z rfers to the variable y 
	px := z.Addr().Interface().(*int) // px := &y
	*px = 3 // y = 3
	fmt.Println(x)

}