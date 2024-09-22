package main

import (
	"time"
	"strconv"
	_"os"
	"fmt"
	"reflect"
	_"io"
)

/*
	Sometimes we need to write a function capable of dealing uniformly with values of types that
	don’t satisfy a common interface, don’t have a known representation, or don’t exist at the time
	we design the function—or even all three.

	Reflection is provided by the reflect package. It defines two important types, Type and
	Value. A Type represents a Go type. It is an interface with many methods for discriminating
	among types and inspecting their components, like the fields of a struct or the parameters of a
	function. The sole implementation of reflect.Type is the type descriptor (§7.5), the same
	entity that identifies the dynamic type of an interface value.

	The inverse operation to reflect.ValueOf is the reflect.Value.Interface method. It
	returns an interface{} holding the same concrete value as the reflect.Value.
*/

// format 

// Any function
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		return strconv.FormatInt(v.Int(),10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,reflect.Uint32, reflect.Uint64,reflect.Uintptr:
		return strconv.FormatUint(v.Uint(),10)
	case reflect.Bool:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func,reflect.Ptr,reflect.Slice,reflect.Map:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()),16)
	default:
		return v.Type().String() + "value"
	}
}


func main() {
	/*
	t := reflect.TypeOf(3) // a reflect.Type
	fmt.Println(t.String())
	fmt.Println(t)
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // os.file
	fmt.Printf("%T\n",3) // int
	v := reflect.ValueOf(5) 
	fmt.Println(v)
	fmt.Printf("%v\n",v)
	fmt.Println(v.String())
	x := v.Type()
	fmt.Println(x.String())

	a := reflect.ValueOf(3)
	b := a.Interface()
	i := b.(int)
	fmt.Printf("%d\n",i)
	*/

	var x int64 = 1 
	var d time.Duration = 1 * time.Nanosecond 
	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any([]int64{x}))
	fmt.Println(Any([]time.Duration{d}))




}