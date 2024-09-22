package main

/*
	In addition to a large collection of the usual mathematical functions, the math package has
	functions for creating and detecting the special values defined by IEEE 754: the positive and
	negative infinities, which represent numbers of excessive magnitude and the result of division
	by zero; and NaN (‘‘not a number’’), the result of such mathematically dubious operations as
	0/0 or Sqrt(-1).
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	The function math.IsNaN tests whether its argument is a not-a-number value, and math.NaN
	returns such a value. It’s tempting to use NaN as a sentinel value in a numeric computation,
	but testing whether a specific computational result is equal to NaN is fraught with peril
	because any comparison with NaN always yields false:
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
*/

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)
	z := f(x,y)
	sx := width / 2 + (x-y) * cos30 * xyscale
	sy := height / 2 + (x+y) * sin30 * xyscale - z * zscale
	return sx, sy
	

}

/*

func graphic() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
	"style='stroke: grey; fill: white; stroke-width: 0.7' " +
	"width='%d' height='%d'>", width, height)

	// Ejercicio 3.3 color rojo para puntas y azul para base
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i + 1 , j)
			bx, by := corner(i,j)
			cx, cy := corner(i, j + 1)
			dx, dy := corner( i + 1, j + 1)
			fmt.Printf("<polyngon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#0000ff;stroke:#ff0000;stroke-width:4'/>\n",ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

*/

func IsInfinite(x float64) bool {
	if math.IsNaN(x) {
		return false
	}else if math.IsInf(x,0) {
		return false
	}
	return true
}

func graphicWeb(writter http.ResponseWriter) {
	writter.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(writter, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	// Ejercicio 3.3 color rojo para puntas y azul para base
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if IsInfinite(ax) && IsInfinite(ay) && IsInfinite(bx) && IsInfinite(by) &&
			   IsInfinite(cx) && IsInfinite(cy) && IsInfinite(dx) && IsInfinite(dy) {
			   fmt.Fprintf(writter, "<polygon points='%g,%g  %g,%g  %g,%g  %g,%g' style='fill:lime;stroke:purple;stroke-width:1'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			}
			   
		}
	}
	fmt.Fprint(writter, "</svg>")
}

func main() {
	//for x := 0; x < 8; x++ {
	//	fmt.Printf("x = %d e**x = %8.3f\n",x,math.Exp(float64(x)))
	//}

	//graphic()

	// Ejercicio 3.3 usar web server
	http.HandleFunc("/show-svg/", func(w http.ResponseWriter, r *http.Request) {
		graphicWeb(w)
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
