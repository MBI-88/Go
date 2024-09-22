package main

/*
	By default, the elements of a new array variable are initially set to
	the zero value for the element type, which is 0 for numbers. We can use an
	array literal to initialize an array with a list of values:
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	In an array literal, if an ellipsis ‘‘...’’ appears in place of the length,
	the array length is determined by the number of initializers. The definition
	of q can be simplified to
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	The size of an array is part of its type, so [3]int and [4]int are different types.
	The size must be a constant expression, that is, an expression whose value can be computed
	as the program is being compiled.
	q := [3]int{1, 2, 3}
	q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int

	In this form, indices can appear in any order and some may be omitted; as before, unspecified
	values take on the zero value for the element type. For instance,
	r := [...]int{99: -1}
	defines an array r with 100 elements, all zero except for the last, which has value −1

	fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int

	Using a pointer to an array is efficient and allows the called function to mutate the caller’s
	variable, but arrays are still inherently inflexible because of their fixed size. The zero function
	will not accept a pointer to a [16]byte variable, for example, nor is there any way to add or
	remove array elements. For these reasons, other than special cases like SHA256’s fixed-size
	hash, arrays are seldom used as function parameters; instead, we use slices


*/

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)
// Currency interface
type Currency int 

const (
	USD Currency = iota
	EUR 
	GBP
	RMB
)

func zero(prt *[32]byte) {
	fmt.Println(*prt)
	*prt = [32]byte{}
	fmt.Println(*prt)
}

func countBit(sh1,sh2 []byte) int {
	lenghtSh1 := len(sh1)
	lenghtSh2 := len(sh2)
	if lenghtSh1 != lenghtSh2 { return -1}

	var count int
	for index := 0; index < lenghtSh1; index++ {
		if sh1[index] != sh2[index] {
			count++
		}
	}
	return count

}

func main(){
	/*
	symbol :=  [...]string{USD:"$",EUR:"E",GBP:"li",RMB:"yu"}
	fmt.Println(RMB,symbol[RMB])

	a := [2]int{1,2}
	b := [...]int{1,2}
	c := [2]int{1,3}
	fmt.Println(a == b, a == c, b == c) // true false false
	
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n",c1,c2,c1 == c2,c1)
	pointer := [32]byte{1,2,4,5}
	zero(&pointer)
	*/

	/* Ejercicio 4.1
	sh1 := sha256.Sum256([]byte("x"))
	sh2 := sha256.Sum256([]byte("X"))
	fmt.Printf("Count different: %d\n",countBit(sh1[:],sh2[:]))
	*/

	 args := os.Args[1:]
	 switch (args[1]){	
	 case ("-384"):
		fmt.Printf("Sha384: %x\n",sha512.Sum384([]byte(args[0])))
		
	 case ("-512"):
		fmt.Printf("Sha512: %x\n",sha512.Sum512([]byte(args[0])))
		
	 default:
		fmt.Printf("Sha256: %x\n",sha256.Sum256([]byte(args[0])))
	 }
	
	


	
}