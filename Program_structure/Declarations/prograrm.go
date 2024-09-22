package main

/*
The constant boilingF is a package-level declaration (as is main), whereas the variables f and
c are local to the function main. The name of each package-level entity is visible not only
throughout the source file that contains its declaration, but throughout all the files of the package.
By contrast, local declarations are visible only within the function in which they are
declared and perhaps only within a small part of it

*/

import "fmt"

const boilingF, freezingF = 212.0, 32.0

func fToC(f float64) float64{
	return (f - 32) * 5 / 9

}


func main(){
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n",boilingF,fToC(boilingF))
}

