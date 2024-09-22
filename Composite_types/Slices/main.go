package main

/*
	Slices represent variable-length sequences whose elements all have the same type. A slice type
	is written []T, where the elements have type T; it looks like an array type without a size

	Arrays and slices are intimately connected. A slice is a lightweight data structure that gives
	access to a subsequence (or perhaps all) of the elements of an array, which is known as the
	slice’s underlying array. A slice has three components: a pointer, a length, and a capacity. The
	pointer points to the first element of the array that is reachable through the slice, which is not
	necessarily the array’s first element. The length is the number of slice elements; it can’t exceed
	the capacity, which is usually the number of elements between the start of the slice and the end
	of the underlying array. The built-in functions len and cap return those values.

	The slice operator s[i:j], where 0 ≤ i ≤ j ≤ cap(s), creates a new slice that refers to elements
	i through j-1 of the sequences, which may be an array variable, a pointer to an array, or
	another slice. The resulting slice has j-i elements. If i is omitted, it’s 0, and if j is omitted, it’s
	len(s). Thus the slice months[1:13] refers to the whole range of valid months, as does the
	slice months[1:]; the slice months[:] refers to the whole array. Let’s define overlapping slices
	for the second quarter and the northern summer.

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2) // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	Slicing beyond cap(s) causes a panic, but slicing beyond len(s) extends the slice, so the
	result may belonger than the original:

	fmt.Println(summer[:20]) // panic: out of range
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer) // "[June July August September October]"

	Unlike arrays, slices are not comparable, so we cannot use == to test whether two slices contain
	the same elements. The standard library provides the highly optimized bytes.Equal function
	for comparing two slices of bytes ([]byte), but for other types of slice, we must do the

	For reference types like pointers and channels, the
	== operator tests reference identity, that is, whether the two entities refer to the same thing. An
	analogous ‘‘shallow’’ equality test for slices could be useful, and it would solve the problem
	with maps, but the inconsistent treatment of slices and arrays by the == operator would be
	confusing. The safest choice is to disallow slice comparisons altogether.

	The only legal slice comparison is against nil, as in
	if summer == nil { ... }

	The zero value of a slice type is nil. A nil slice has no underlying array. The nil slice has
	length and capacity zero, but there are also non-nil slices of length and capacity zero, such as
	[]int{} or make([]int, 3)[3:]. As with any type that can have nil values, the nil value of a
	particular slice type can be written using a conversion expression such as []int(nil).
	var s []int // len(s) == 0, s == nil
	s = nil // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{} // len(s) == 0, s != nil

	The built-in function make creates a slice of a specified element type, length, and capacity. The
	capacity argument may be omitted, in which case the capacity equals the length.
	make([]T, len)
	make([]T, len, cap) // same as make([]T, cap)[:len]

	The append Function

	Updating the slice variable is required not just when calling append, but for any function that
	may change the length or capacity of a slice or make it refer to a different underlying array. To
	use slices correctly, it’s important to bear in mind that although the elements of the underlying
	array are indirect, the slice’s pointer, length, and capacity are not. To update them requires an
	assignment like the one above. In this respect, slices are not ‘‘pure’’ reference types but resemble an aggregate type 
	such as this struct:
	
	type IntSlice struct {
		ptr *int
		len, cap int
	}

	Our appendInt function adds a single element to a slice, but the built-in append lets us add
	more than one new element, or even a whole slice of them.

	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println(x) // "[1 2 3 4 5 6 1 2 3 4 5 6]"

*/

import (
	//"crypto/x509"
	"fmt"
	"bytes"
	"unicode/utf8"
	"unicode"
)

func reverse(s []int) {
	for i,j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i],s[j] = s[j], s[i]
	}
}

func equal(x,y []string) bool {
	if len(x) != len(y) { return false}
	for i := range x {
		if x[i] != y[i] { return false}
	}
	return true
}

// Simulando el proceso de append
func appendInt(x []int,y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	}else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z,x)
	}
	z[len(x)] = y
	return z
}

func remove(slice []int, i int) []int {
	copy(slice[i:],slice[i+1:])
	return slice[:len(slice) - 1]
}

func reversePoint(s *[6]int) {
	for i,j := 0, len(s) - 1; i < j; i,j = i+1,j-1 {
		s[i],s[j] = s[j],s[i]

	}
}

func rotate(s []int, pos int) []int {
	if pos < len(s) {
		n := s[:pos]
		s = append(s[pos:], n...)
		return s
	}else {
		return s
	}
}

func duplicated(s []string) []string {
	for i := 0; i < len(s) - 1; i++ {
		if s[i] == s[len(s) + i - len(s) + 1]{
			n := s[:i]
			n = append(n,s[i+1:]...)
			s = n
		}else{
			continue
		}
	}
	return s
}

func removesquashes(b []byte) []byte {
	var buff bytes.Buffer

	for i := 0; i < len(b); i++ {
			r,size := utf8.DecodeRuneInString(string(b[i]))
			if unicode.IsSpace(r) {
				nspace,_ := utf8.DecodeLastRuneInString(string(b[i+size:]))
				if unicode.IsSpace(nspace){
					buff.WriteRune(' ')
				}
			}else{
				buff.WriteRune(r)
			}
	}


	return buff.Bytes()
}

func reverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reversebyte(b[i: i + size])
		i += size
	}
	reversebyte(b)
	return b
}

func reversebyte(b []byte) {
	for i,j := 0, len(b) - 1; i < j; i,j = i+1,j-1 {
		b[i],b[j] = b[j],b[i]
	}

}

func main() {
	/*
	valor := []int{0,1,2,3,4,5,6,7,8,9}
	reverse(valor[:])
	fmt.Printf("Reverse: %d\n",valor[:])

	fmt.Printf("Equals: %t\n",equal([]string{"a","b","c"},[]string{"a","b","c"}))
	
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x,i)
		fmt.Printf("%d cap=%d\t%v\n",i,cap(y), y)
		x = y
	}

	s := []int{5,6,7,8,9}
	fmt.Println(remove(s,2))
	*/
	// Ejercicio 4.3
	//array := [...]int{1,2,3,4,5,6}
	//reversePoint(&array)
	//fmt.Printf("Reverse: %d\n",array)

	// Ejercicio 4.4
	//array := []int{1,2,3,4,5}
	//fmt.Printf("Rotate %d %d\n",2,rotate(array,3))

	// Ejercicio 4.5
	//stringArray := [...]string{"H","o","l","a"," ","m","u","u","n","d","o"}
	//fmt.Printf("Array %s  New array %s\n",stringArray,duplicated(stringArray[:]))

	// Ejercicio 4.6
	//words := "1  +  2    =  3"
	//fmt.Printf("Remove spaces %s --> %s\n",words,removesquashes([]byte(words)))

	// Ejercicio 4.7
	valor := "Lorem ipsum"
	fmt.Printf("Reverse %s --> %s\n",valor,reverseUTF8([]byte(valor)))



}