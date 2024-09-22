package main

/*
	The Constant Generator iota

	A const declaration may use the constant generator iota, which is used to create a sequence
	of related values without spelling out each one explicitly. In a const declaration, the value of
	iota begins at zero and increments by one for each item in the sequence.

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	Untyped Constants

	For literals, syntax determines flavor. The literals 0, 0.0, 0i, and '\u0000' all denote const ants of the same value but different
	flavors: untyped integer, untyped floating-point, untyped
    complex, and untyped rune, respectively. Similarly, true and false are untyped booleans and
    string literals are untyped strings.

	In a variable declaration without an explicit type (including short variable declarations), the
	flavor of the untyped constant implicitly determines the default type of the variable, as in these
	examples:
	i := 0 // untyped integer; implicit int(0)
	r := '\000' // untyped rune; implicit rune('\000')
	f := 0.0 // untyped floating-point; implicit float64(0.0)
	c := 0i // untyped complex; implicit complex128(0i)


*/

import (
	"fmt"
)

type Flags uint 

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPoinToPoint
	FlagMulticast
)

func IsUp(v Flags) bool {
	return v & FlagUp == FlagUp
}
func TurnDown(v *Flags) { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast}
func IsCast(v Flags) bool {
	return v & (FlagBroadcast | FlagMulticast) != 0
}

const (
	KB = 1000 
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB 
	EB = PB * KB 
	ZB = EB * KB 
	YB = ZB * KB
)



func main() {
	/*
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n",v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n",v,IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n",v,IsUp(v))
	fmt.Printf("%b %t\n",v,IsCast(v))
	*/
	// Ejercicio 3.13

	fmt.Printf("%d %d %d\n",KB,MB,GB)
}
