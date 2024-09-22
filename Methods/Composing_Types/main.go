package main

/*
	Methods can be declared only on named types (like Point) and pointers to them (*Point),
	but thanks to embedding, it’s possible and sometimes useful for unnamed struct types to have
	methods too.

*/

import (
	"fmt"
	"image/color"
	"math"
)

// coloredpoint

// Point struct
type Point struct {
	X,Y float64
}

// Distance method of Point
func (p *Point) Distance(q *Point) float64 {
	return math.Hypot((q.X - p.X), (q.Y - p.Y))
}

// ColoredPoint point struct
type ColoredPoint struct {
	Point 
	Color color.RGBA
}



func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2 
	fmt.Println(cp.Y) // "2"

	red := color.RGBA{255,0,0,255}
	blue := color.RGBA{0,0,255,255}
	var p = ColoredPoint{Point{1,1},red}
	var q = ColoredPoint{Point{5,4},blue}
	fmt.Println(p.Distance(&q.Point)) // "5"



}