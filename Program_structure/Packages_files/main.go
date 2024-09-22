package main

import (
	"fmt"
	"os"
	"packages/tempconv"
	"strconv"
)

var fval tempconv.Fahrenheit = 200
var kval tempconv.Kelvin = 100

func converUnit() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg,64)
		if err != nil {
			fmt.Fprintf(os.Stderr,"cf: %v\n",err)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		meter := tempconv.Meters(t)
		pound := tempconv.Pounds(t)
		feet := tempconv.Feet(t)
		kilograms := tempconv.Kilograms(t)
		kelvin := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s\n",f,tempconv.FToC(f),c,tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n",meter,tempconv.MeToFe(meter),feet,tempconv.FeToMe(feet))
		fmt.Printf("%s = %s, %s = %s\n",pound,tempconv.PToKil(pound),kilograms,tempconv.KilToP(kilograms))
		fmt.Printf("%s = %s, %s = %s\n",kelvin,tempconv.KToC(kelvin),kelvin,tempconv.KToF(kelvin))
	}
}

func main(){
	//fmt.Printf("Probando imports %g\n",tempconv.AbsoluteZeroC)
	//fmt.Printf("F to C => %g result => %s\n",fval,tempconv.FToC(fval))
	//fmt.Printf("K to F => %g result => %s\n",kval,tempconv.KToF(kval))

	// Ejercicio 2.2

	converUnit()
}

