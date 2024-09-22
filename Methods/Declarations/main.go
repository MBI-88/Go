package main

/*
	Although there is no universally accepted definition of object-oriented programming, for our
	purposes, an object is simply a value or variable that has methods, and a method is a function
	associated with a particular type. An object-oriented program is one that uses methods to
	express the properties and operations of each data structure so that clients need not access the
	object’s representation directly.

	A method is declared with a variant of the ordinary function declaration in which an extra
	parameter appears before the function name. The parameter attaches the function to the type
	of that parameter.

	In Go, we don’t use a special name like this or self for the receiver; we choose receiver
	names just as we would for any other parameter. Since the receiver name will be fre quently
	us ed, it’s a good ide a to choose something short and to be con sistent across met hods. 
	A common choice is the firs t letter of the typ e name, like p for Point.

	The expression p.Distance is called a selector, because it selects the appropriate Distance
	method for the receiver p of type Point. Selectors are also used to select fields of struct types,
	as in p.X. Since methods and fields inhabit the same name space, declaring a method X on the
	struct type Point would be ambiguous and the compiler will reject it.

	All methods of a given type must have unique names, but different types can use the same
	name for a method, like the Distance methods for Point and Path; there’s no need to qualify
	function names (for example, PathDistance) to disambiguate. Here we see the first benefit to
	using methods over ordinary functions: method names can be shorter. The benefit is magnified 
	for calls originating outside the package, since they can use the shorter name and omit the
	package name.
	
	
*/


import (
	"fmt"
	"math"
)

// Point struct 
type Point struct {
	X,Y float64
}

// Path struct
type Path []Point

// Distance method of Point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

// Distance method of Path
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum  += path[i - 1].Distance(path[i])
		}
	}
	return sum
}


func main() {
	/*
	p := Point{1,2}
	q := Point{4,6}

	fmt.Println(p.Distance(q))
	*/
	perim := Path{
		{1,2},
		{5,1},
		{5,4},
		{1,1},
	}

	fmt.Println(perim.Distance()) // 12


}