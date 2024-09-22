package main

/*
	Because calling a function makes a copy of each argument value, if a function needs to update
	a variable, or if an argument is so large that we wish to avoid copying it, we must pass the
	address of the variable using a pointer. The same goes for methods that need to update the
	receiver variable: we attach them to the pointer type, such as *Point.

	func (p *Point) ScaleBy(factor float64) {
		p.X *= factor
		p.Y *= factor
	}

	The name of this method is (*Point).ScaleBy. The parentheses are necessary ; without
	them, the expression would be parsed as *(Point.ScaleBy).

	In a realistic program, convention dictates that if any method of Point has a pointer receiver,
	then all methods of Point should have a pointer receiver, even ones that don’t strictly need it.
	We’ve broken this rule for Point so that we can show both kinds of method.

	Named types (Point) and pointers to them (*Point) are the only types that may appear in a
	receiver declaration. Furthermore, to avoid ambiguities, method declarations are not permitted 
	on named types that are themselves pointer types:

		type P *int

		func (P) f() { ... } // compile error: invalid receiver type

	The (*Point).ScaleBy method can be cal le d by providing a *Point re ceiver, like this:
		r := &Point{1, 2}
		r.ScaleBy(2)
		fmt.Println(*r) // "{2, 4}"
	or this:
		p := Point{1, 2}
		pptr := &p
		pptr.ScaleBy(2)
		fmt.Println(p) // "{2, 4}"
	or this:
		p := Point{1, 2}
		(&p).ScaleBy(2)
		fmt.Println(p) // "{2, 4}"
	
	But the last two cases are ungainly. Fortunately, the language helps us here. If the receiver p is
	a variable of type Point but the method requires a *Pointre ceiver, we can use this shorthand

	p.ScaleBy(2)

	and the compiler will perform an implicit &p on the variable. This works only for variables,
	including struct fields like p.X and array or slice elements like perim[0]. We cannot call a
	*Point method on a non-addressable Point receiver, because there’s no way to obtain the
	address of temporary value.

	Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal

	But we can call a Point method like Point.Distance with a *Point receiver, because there is
	a way to obtain the value from the address: just load the value pointed to by the receiver. The
	compiler inserts an implicit * operation for us. These two function calls are equivalent:

	pptr.Distance(q)
	(*pptr).Distance(q)

	If all the methods of a named type T have a receiver type of T itself (not *T), it is safe to copy
	instances of that type; calling any of its methods necessarily makes a copy. For example,
	time.Duration values are liberally copied, including as arguments to functions. But if any
	method has a pointer receiver, you should avoid copying instances of T because doing so may
	violate internal invariants. For example, copying an instance of bytes.Buffer would cause
	the original and the copy to alias (§2.3.2) the same underlying array of bytes. Subsequent
	method calls would have unpredictable effects.

	Nil is a valid recevier value

	Just as some functions allow nil pointers as arguments, so do some methods for their receiver,
	especially if nil is a meaningful zero value of the type, as with maps and slices. In this simple
	linked list of integers, nil represents the empty list

*/


import (
	"fmt"
)

// IntList is a linked list  
type IntList struct {
	Value int
	Tail *IntList
}

// Sum method of IntList
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

// net/url

// Values map
type Values map[string][]string

// Get method of Values
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add method of Get
func (v Values) Add(key,value string) {
	v[key] = append(v[key],value)
}

func main() {
	/*
	linked := IntList{Value:2,Tail:nil}
	nextlist := IntList{Value:5,Tail: &linked}
	fmt.Printf("Sum: %d\n",nextlist.Sum())
	*/

	m := Values{"lang": {"en"},}
	m.Add("item","1")
	m.Add("item","2")
	fmt.Printf("Language: %s\n",m.Get("lang"))
	fmt.Printf("Query: %s\n",m.Get("q"))
	fmt.Printf("First value: %s\n", m.Get("item"))
	fmt.Println(m["item"])

}