package main

/*
	The type of a variable or expression defines the characteristics of the values it may take on,
	such as their size (number of bits or number of elements, perhaps), how they are represented
	internally, the intrinsic operations that can be performed on them, and the methods associated 
	with them.

	A type declaration defines a new named type that has the same underlying type as an existing
	type. The named type provides a way to separate different and perhaps incompatible uses of
	the underlying type so that they can’t be mixed unintentionally.

	type name underlying-type

	Type declarations most often appear at package level, where the named type is visible throughout the package, 
	and if the name is exported (it starts wit h an upper-case letter), it’s accessible
	from other packages as well.	

	This package defines two types, Celsius and Fahrenheit, for the two units of temperature.
	Even though both have the same underlying type, float64, they are not the same type, so they
	cannot be compared or combined in arithmetic expressions. Distinguishing the types makes it
	possible to avoid errors like inadvertently combining temperatures in the two different scales;
	an explicit type conversion like Celsius(t) or Fahrenheit(t) is required to convert from a
	float64. Celsius(t) and Fahrenheit(t) are conversions, not function calls. They don’t
	change the value or represent ation in any way, but they make the change of meaning explicit.
	On the other hand, the functions CToF and FToC convert between the two scales; they do
	return different values.

	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	fmt.Printf("%g\n", boilingF-FreezingC) // compile error: type mismatch

	Comparison operators like == and < can also be used to compare a value of a named type to
	another of the same type, or to a value of the underlying type. But two values of different
	named types cannot be compared directly :

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	fmt.Println(c == f) // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true"!

	The declaration below, in which the Celsius parameter c appears before the function name,
	associates with the Celsius type a method named String that returns c’s numeric value
	followe d by °C:
*/

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32)* 5/9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C",c)
}

func main(){
	// Uso de la centencia type para declarar un nuevo tipo
	// de dato
	//fmt.Printf("CToF => %g\n FToC => %g\n",CToF(25),FToC(220))
	c := FToC(212.0)
	fmt.Println(c.String()) // 100°C
	fmt.Printf("%v\n",c) // 100°C
	fmt.Printf("%s\n",c) // 100°C
	fmt.Println(c) // 100°C
	fmt.Printf("%g\n",c)
	fmt.Println(float64(c)) // 100


}