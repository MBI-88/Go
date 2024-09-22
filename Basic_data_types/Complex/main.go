package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}

		}

	}
	return color.Black
}

var paleta = []color.Color{
	color.RGBA{0xf4, 0x43, 0x36, 0xff}, 
	color.RGBA{0xff, 0xff, 0xff, 0xff}, 
	color.RGBA{0x4c, 0xaf, 0x50, 0xff}, 
	color.RGBA{0xff, 0xff, 0xff, 0xff}, 
	color.RGBA{0x21, 0x96, 0xf3, 0xff}, 
	color.RGBA{0xff, 0xff, 0xff, 0xff}, 
}

func mandelbrotColor(z complex128) color.Color {
	const iterations = 200
	var v complex128

	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return paleta[n % len(paleta)]
		}
	}
	return color.Black
}

func AverageColor(colors []color.Color) color.Color {
	if len(colors) < 1 {
		return nil
	}
	var r,g,b,a  float64
	for _, color := range colors {
		dr, dg, db, da := color.RGBA()
		r += float64(dr >> 8) / float64(len(colors))
		g += float64(dg >> 8) / float64(len(colors))
		b += float64(db >> 8) / float64(len(colors))
		a += float64(da >> 8) / float64(len(colors))
	} 

	return color.RGBA{uint8(r),uint8(g),uint8(b),uint8(a)}
}

func equalsFloat64(x,y float64) bool {
	return math.Abs(x - y) < 1e-6

}

func NewtonMethod(z complex128) color.Color{
	// f(x) = z^4 - 1
	// z' = z - f(z)/f'(z)
	//    = z - (z^4 - 1) / (4 * z^3)
	//    = z - (z - 1/z^3) / 4
	const iterations = 37
	const contrast = 7
	for n := uint8(0); n < iterations; n++ {
		z -= (z-1/(z*z*z))/4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			if equalsFloat64(real(z), 1) && equalsFloat64(imag(z), 0) {
				return color.RGBA{255 - contrast*n, 0, 0, 0xff}
			} else if equalsFloat64(real(z), -1) && equalsFloat64(imag(z), 0) {
				return color.RGBA{0, 255 - contrast*n, 0, 0xff}
			} else if equalsFloat64(real(z), 0) && equalsFloat64(imag(z), 1) {
				return color.RGBA{0, 0, 255 - contrast*n, 0xff}
			}
			return color.RGBA{255 - contrast*n, 255 - contrast*n, 0, 0xff}
		}
	}
	return color.Black
}

func main() {

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		//y0 := float64(py) / height * (ymax - ymin) + ymin //Ejercicio 3.6
		//y1 := (float64(py) + 0.25) / height *  (ymax - ymin) + ymin //Ejercicio 3.6
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin
			z := complex(x, y) 
			img.Set(px, py, NewtonMethod(z))
			//img.Set(px,py,mandelbrotColor(z)) //Ejercicio 3.5

			/* Ejercicio 3.6
			x0 := float64(px) / width * (xmax - xmin) + xmin
			x1 := (float64(px) + 0.25) / width * (xmax - xmin) + xmin
			z0 := complex(x0,y0)
			z1 := complex(x1,y0)
			z2 := complex(x0,y1)
			z3 := complex(x1, y1)
			colorSelected := AverageColor([]color.Color{
				mandelbrotColor(z0),
				mandelbrotColor(z1),
				mandelbrotColor(z2),
				mandelbrotColor(z3),
			})
			img.Set(px,py,colorSelected) 
			*/


		}
	}



	//f,_ := os.Create("mandelbrot.png")
	
	//f, _ := os.Create("mandelbrot_3.5.png") // Ejercicio 3.5
	
	//f, _ := os.Create("mandelbrot_3.6.png") // Ejercicio 3.6
	f,_ := os.Create("Newton_factral.png") // Ejercicio 3.7 usa las variables basicas de los bucles for

	png.Encode(f, img)
}
