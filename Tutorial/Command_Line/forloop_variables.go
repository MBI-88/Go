package main

import (
	"fmt"
	"os"
	//"strings"

	
)


func Forloop() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func ForloopRange() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	}

	fmt.Println(s)
}


func exercise1() string{
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func exercise2() {
	for index,value := range os.Args {
		fmt.Println(index)
		fmt.Println(value)
	}
}


func main() {
	// Demostrando el uso de for y la asignacion de variables en sus diversas formas
	//fmt.Printf("Hola mundo de Golang estoy comenzando")
	//Forloop()
	//ForloopRange()
	// Version eficiente
	//fmt.Println(strings.Join(os.Args[1:]," "))

	// Ejercicio 1

	//result :=  exercise1()
	// version optima
	//fmt.Println(strings.Join(os.Args[0:])," ")
	//fmt.Println(result)


	// Ejercicio 2
	exercise2()

}
