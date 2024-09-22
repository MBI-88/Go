package main 

import (
	"unsafe"
	_"fmt"
)

/*
	The unsafe package is rather magical. Although it appears to be a regular package and is
	imported in the usual way, it is actually implemented by the compiler. It provides access to a
	number of built-in language features that are not ordinarily available because they expose
	details of Go’s memory layout. Presenting these features as a separate package makes the rare
	occasions on which they are needed more conspicuous. Also, some environments may restrict
	the use of the unsafe package for security reasons.

	The unsafe.Sizeof function reports the size in bytes of the representation of its operand,
	which may be an expression of any type; the expression is not evaluated. A call to Sizeof is a
	constant expression of type uintptr, so the result may be used as the dimension of an array
	type, or to compute other constants.

	import "unsafe"
	fmt.Println(unsafe.Sizeof(float64(0))) // "8"

	Type                          Size
	bool                          1 byte
	intN, uintN, floatN, complexN N /8bytes (for example, float64 is 8 bytes)
	int, uint, uintptr            1 word
	*T                            1 word
	string                        2 words (data, len)
	[]T                           3 words (data, len, cap)
	map                           1 word
	func 						  1 word
	chan                          1 word
	interface                     2 words (type, value)


	The unsafe.Alignof function reports the required alignment of its argument’s type. Like
	Sizeof, it may be applied to an expression of any type, and it yields a constant. Typically,
	boolean and numeric types are aligned to their size (up to a maximum of 8 bytes) and all other
	types are word-aligned.

	unsafe.Pointer

	An ordinary *T pointer may be converted to an unsafe.Pointer, and an unsafe.Pointer
	may be converted back to an ordinary pointer, not necessarily of the same type *T. 
	By converting a *float64 pointer to a *uint64, for instance, we can inspect the bit pattern of a 
	floating-point variable.

	The reason is very subtle. Some garbage collectors move variables around in memory to
	reduce fragmentation or bookkeeping. Garbage collectors of this kind are known as moving
	GCs. When a variable is moved, all pointers that hold the address of the old location must be
	updated to point to the new one. From the perspective of the garbage collector, an
	unsafe.Pointer is a pointer and thus its value must change as the variable moves, but a
	uintptr is just a number so its value must not change . The incorrect code above hides a
	pointer from the garbage collector in the non-pointer variable tmp. By the time the second
	statement executes, the variable x could have moved and the number in tmp would no longer
	be the address &x.b. The third statement clobbers an arbitrary memory location with the
	value 42.

	At the time of writing , there is little clear guidance on what Go programmers may rely upon
	after an unsafe.Pointer to uintptr conversion (see Go issue 7192), so we strongly recommend that you assume 
	the bare minimum. Treat all uintptr values as if they contain the
	former address of a variable, and minimize the number of operations between converting an
	unsafe.Pointer to a uintptr and using that uintptr. In our first example above , the three
	operations—conversion to a uintptr, addition of the field offset, conversion back—all
	appeared within a single expression.

*/

// Float64bit function
func Float64bit(f float64) uint64 {return *(*uint64)(unsafe.Pointer(&f))}

// unsafeptr 

var x struct {
	a bool 
	b int16 
	c []int
}

func main() {
	/*
	var x struct {
		a bool
		b int16 
		c []int
	}

	fmt.Printf("%v\n %v\n %v\n %v\n",
	unsafe.Sizeof(x),unsafe.Sizeof(x.a),unsafe.Sizeof(x.b),
	unsafe.Sizeof(x.c))

	fmt.Printf("%#016x\n",Float64bit(1.0))

	*/
	// equivalente a pb := &x.b
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42






}