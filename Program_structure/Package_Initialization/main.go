package main

/*
	Package initialization begins by initializing package-level variables in the order in which they
	are declared, except that dependencies are resolved first.

	If the package has multiple .go files, they are initialized in the order in which the files are
	given to the compiler ; the go tool sorts .go files by name before invoking the compiler.

	Each variable declared at package level starts life with the value of its initializer expression, if
	any, but for some variables, like tables of data, an initializer expression may not be the simplest
	way to set its initial value. In that case, the init function mechanism may be simpler. Any
	file may contain any number of functions whose declaration is just

	func init(){...}

	Such init functions can’t be called or referenced, but otherwise they are normal functions.
	Within each file, init functions are automatically executed when the program starts, in the
	order in which they are declared.

	One package is initialized at a time, in the order of imports in the program, dependencies first,
	so a package p importing q can be sure that q is fully initialized before p’s initialization begins.
	Initialization proceeds from the bottom up; the main package is the last to be initialized. In
	this manner, all packages are fully initialized before the application’s main function begins
*/

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pc [256]byte

func init(){
	for i := range pc {
		pc[i] = pc[i/2] + byte(i & 1)
	}
}

func PopCount(x uint64) int {
	return int(
		pc[byte(x >> (0*8))] +
		pc[byte(x >> (1*8))] +
		pc[byte(x >> (2*8))] +
		pc[byte(x >> (3*8))] +
		pc[byte(x >> (4*8))] +
		pc[byte(x >> (5*8))] +
		pc[byte(x >> (6*8))] +
		pc[byte(x >> (7*8))])
}

func PopCountLoop(x uint64) int {
	var result int
	for i := range pc {
		result += int(pc[byte(x >> (i*8))]) 
	}
	return result
}

func PopCount64(x uint64) int {
	var result int 
	for i := range pc {
		result += int(pc[byte(x >> (i*64))])
	}
	return result
}

func PopCountNonZero(x uint64) int {
	var result int
	for i := range pc {
		result += int(pc[byte( x & (x-1) >> (i*8))])
	}
	return result
}

func main(){
	args := os.Args[1]
	args = strings.Split(args," ")[0]
	result, err := strconv.ParseUint(args,2,64)
	if err != nil {
		//fmt.Printf("%d\n",PopCount(result))

		// Ejercicio 2.3
		//fmt.Printf("%d\n",PopCountLoop(result))

		// Ejercicio 2.4
		//fmt.Printf("%d\n",PopCount64(result))

		// Ejercicio 2.5
		fmt.Printf("%d\n",PopCountNonZero(result))


	}else {
		fmt.Println("Error convertion")
	}

	


	
}