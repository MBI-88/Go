package main

import (
	"fmt"
	"reflect"
	"unsafe"
	"math"
)

/*
	The DeepEqual function from the reflect package reports whether two values are ‘‘deeply’’
	equal. DeepEqual compares basic values as if by the built-in == operator ; for composite values,
	it traverses them recursively, comparing corresponding elements. Because it works for
	any pair of values, even ones that are not comparable with ==, it finds widespread use in tests.
	The following test uses DeepEqual to compare two []string values

*/

// equal

const epsilon = 0.0000001

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()
	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)
	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true
	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !equal(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true

	case reflect.Int, reflect.Int32, reflect.Int64: // Ejercicio 13.1
		return x.Int() == y.Int()

	case reflect.Float32, reflect.Float64: // Ejercicio 13.1
		return math.Abs(x.Float() - y.Float()) < epsilon

	case reflect.String:
		return x.String() == y.String()

	}
	panic("unreachable")
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

// Equal function
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

func main() {
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(Equal([]string{"foo"}, []string{"bar"}))
	fmt.Println(Equal(map[string]int(nil), map[string]int{}))
	fmt.Println(Equal(1000,1005))

}
