package main

import (
	"os"
)

/*
	A struct is an aggregate data type that groups together zero or more named values of arbitrary
	types as a single entity. Each value is called a field. The classic example of a struct from data
	processing is the employee record, whose fields are a unique ID, the employee’s name, address,
	date of birth, position, salary, manager, and the like. All of these fields are collected into a single entity
	that can be copied as a unit, passed to functions and returned by them, stored in
	arrays, and so on.

	These two statements declare a struct type called Employee and a variable called dilbert that
	is an instance of an Employee:

	type Employee struct {
		ID int
		Name string
		Address string
		DoB time.Time
		Position string
		Salary int
		ManagerID int
	}

	var dilbert Employee

	The individual fields of dilbert are accessed using dot notation like dilbert.Name and
	dilbert.DoB. Because dilbert is a variable, its fields are variables too, so we may assign to a
	field:
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code

	or take its address and access it through a pointer:
	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	The dot notation also works with a pointer to a struct:

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	The last statement is equivalent to

	(*employeeOfTheMonth).Position += " (proactive team player)"

	Given an employee’s unique ID, the function EmployeeByID returns a pointer to an Employee
	struct. We can use the dot notation to access its fields:

	func EmployeeByID(id int) *Employee { ...  }
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for... no real reason

	The last statement updates the Employee struct that is pointed to by the result of the call to
	EmployeeByID. If the result type of EmployeeByID were changed to Employee instead of
	*Employee, the assignment statement would not compile since its left-hand side would not
	identify a variable.

	The name of a struct field is exported if it begins with a capital letter; this is Go’s main access
	control mechanism. A struct type may contain a mixture of exported and unexported fields.

	A named struct type S can’t declare a field of the same type S: an aggregate value cannot contain it's self.
	(An analogous restriction applies to arrays.) But S may declare a field of the
	pointer type *S, which lets us create recursive data structures like linked lists and trees. This is
	illustrated in the code below, which uses a binary tree to implement an insertion sort.

	The struct type with no fields is called the empty struct, written struct{}. It has size zero and
	carries no information but may be useful nonetheless. Some Go programmers use it instead
	of bool as the value type of a map that represents a set, to emphasize that only the keys are significant,
	but the space saving is marginal and the syntax more cumbersome, so we generally avoid it.

	seen := make(map[string]struct{}) // set of strings
	// ...
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
		// ...first time seeing s...
	}

	Struct Literals

	type Point struct{ X, Y int }
	p := Point{1, 2}

	There are two forms of struct literal. The first form, shown above , requires that a value be
	specified for every field, in the right order. It burdens the writer (and reader) with remembering exactly
	what the fields are , and it makes the code fragile should the set of fields later grow
	or be reordered. Accordingly, this form tends to be used only within the package that defines
	the struct type, or with smaller struct types for which there is an obvious field ordering convention,
	like image.Point{x, y} or color.RGBA{red, green, blue, alpha}

	More often, the second form is used, in which a struct value is initialized by listing some or all
	of the field names and their corresponding values, as in this statement from the Lissajous
	program of Section 1.4.
	anim := gif.GIF{LoopCount: nframes}

	The two forms cannot be mixed in the same literal. Nor can you use the (order-based) first
	form of literal to sneak around the rule that unexported identifiers may not be referred to
	from another package.

	package p
	type T struct{ a, b int } // a and b are not exported

	package q

	import "p"
	var _ = p.T{a: 1, b: 2} // compile error: can't reference a, b
	var _ = p.T{1, 2} // compile error: can't reference a, b

	For efficiency, larger struct types are usually passed to or returned from functions indirectly
	using a pointer,

	func Bonus(e *Employee, percent int) int {
		return e.Salary * percent / 100
	}

	and this is required if the function must modify its argument, since in a call-by-value language
	like Go, the called function receives only a copy of an argument, not a reference to the original
	argument.

	func AwardAnnualRaise(e *Employee) {
		e.Salary = e.Salary * 105 / 100
	}

	Because structs are so commonly dealt with through pointers, it’s possible to use this
	shorthand notation to create and initialize a struct variable and obtain its address:

	pp := &Point{1, 2}

	It is exactly equivalent to:

	pp := new(Point)
	*pp = Point{1, 2}

	but &Point{1,2} can be used directly within an expression, such as a function call.


	Comparing Structs

	If all the fields of a struct are comparable, the struct itself is comparable, so two expressions of
	that type may be compared using == or !=. The == operation compares the corresponding
	fields of the two structs in order, so the two printed expressions below are equivalent:

	type Point struct{ X, Y int }
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q) // "false"

	Comparable struct types, like other comparable types, may be used as the key type of a map.
	type address struct {
		hostname string
		port int
	}
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++


	Struct Embedding and Anonymous Fields

	Consider a 2-D drawing program that provides a library of shapes, such as rectangles, ellipses,
	stars, and wheels. Here are two of the types it might define:

	type Circle struct {
		X, Y, Radius int
	}
	type Wheel struct {
		X, Y, Radius, Spokes int
	}

	As the set of shapes grows, we’re bound to notice similarities and repetition among them, so it
	may be convenient to factor out their common parts:

	type Point struct {
		X, Y int
	}
	type Circle struct {
		Center Point
		Radius int
	}
	type Wheel struct {
		Circle Circle
		Spokes int
	}

	The application maybe clearer for it, but this change makes accessing the fields of a Wheel
	more verbose:

	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

	Go lets us declare a field with a type but no name; such fields are called anonymous fields. The
	type of the field must be a named type or a pointer to a named type. Below, Circle and Wheel
	have one anonymous field each. We say that a Point is embedded within Circle, and a
	Circle is embedded within Wheel.

	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes int
	}

	var w Wheel
	w.X = 8 // equivalent to w.Circle.Point.X = 8
	w.Y = 8 // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20

	Because ‘‘anonymous’’ fields do have implicit names, you can’t have two anonymous fields of
	the same type since their names would conflict. And because the name of the field is implicitly 
	determined by its type, so too is the visibility of the field. In the examples above, the Point
	and Circle anonymous fields are exported. Had they been unexported (point and circle),
	we could still use the shorthand form:

	w.X = 8 // equivalent to w.circle.point.X = 8

	but the explicit long form shown in the comment would be forbidden outside the declaring
	package because circle and point would be inaccessible.

*/

import (
	"fmt"
)

// Tree sort

type tree struct {
	value       int
	left, right *tree
}

func sort(t *tree,values []int) {
	for _, v := range values {
		t = add(t, v)
	}
	appendValues(values[:0], t)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func printTree(t *tree) {

	if t == nil {
		fmt.Println("Tree empty")
		os.Exit(0)
	}
	fmt.Printf("Node value: %v\n",t.value)
	if t.left != nil {
		printTree(t.left)
	}
	if t.right != nil {
		printTree(t.right)
	}

}

/*Point struct */
type Point struct {
	X, Y int
}

/*Circle struct */
type Circle struct {
	Point  
	Radius int
}

/*Wheel struct */
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	/*
	instance := Wheel{Circle{Point{8, 8}, 5}, 20}

	fmt.Printf("Struct Wheel: %#v\n", instance)
	*/
	root := new(tree)
	sort(root,[]int{1,2,0,1,2,3})
	printTree(root)


}
