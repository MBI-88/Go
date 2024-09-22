package main

/*
	Usually we select and call a method in the same expression, as in p.Distance(), but it’s possible to separate these two operations. The selector p.Distance yields a method value,
	a function that binds a method (Point.Distance) to a specific receiver value p. This function can
	then be invoked without a receiver value; it needs only the non-receiver arguments.
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance // method value
	fmt.Println(distanceFromP(q)) // "5"
	var origin Point // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", ;5
	scaleP := p.ScaleBy // method value
	scaleP(2) // p becomes (2, 4)
	scaleP(3) // then (6, 12)
	scaleP(10) // then (60, 120)

	Method expressions can be helpful when you need a value to represent a choice among several
	methods belonging to the same type so that you can call the chosen method with many
	different receivers. In the following example, the variable op represents either the addition or
	the subtraction method of type Point, and Path.TranslateBy calls it foreach pointin the
	Path.



*/

import (
	"fmt"
)

// Point struct
type Point struct {
	X,Y float64
}

// Add method of Point
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Sub method of Point
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}


// Path of point struct
type Path []Point

// TranslateBy method of Path
func (path Path) TranslateBy(offset Point, add bool) Path {
	var op func(p,q Point) Point
	if add {
		op = Point.Add

	}else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i],offset)
	}
	return path
}



func main() {
	p  := Point{5,5}
	q := Point{10,10}
	path := Path{p,q}

	fmt.Printf("%v\n",path.TranslateBy(p,true))


}