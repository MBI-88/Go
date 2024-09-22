package main

import (
	"strings"
	"fmt"
	"reflect"
	"strconv"
)

/*
	Slices and arrays: The logic is the same for both. The Len method returns the number of elements of a slice or array value,
	and Index(i) retrieves the element at index i, also as a
	reflect.Value; it panics if i is out of bounds. These are analogous to the built-in len(a)
	and a[i] operations on sequences. The display function recursively invokes itself on each
	element of the sequence, app ending the subscript not ation "[i]" to the path.

	Structs: The NumField method reports the number of fields in the struct, and Field(i)
	returns the value of the i-th field as a reflect.Value. The list of fields includes ones
	promoted from anonymous fields. To app end the field selector notation ".f" to the path, we
	must obtain the reflect.Type of the struct and access the name of its i-th field.

	Maps: The MapKeys method return saslice of reflect.Values, one per map key. As usual
	when iterating overamap, the order is undefined. MapIndex(key) returns the value corresponding to key.
	We append the subscript not ation "[key]" to the path. (We’re cutting a
	corner here. The type of a map key isn’t restricted to the types formatAtom handles best;
	arrays, structs, and interfaces can also be valid map keys. Extending this case to print the key
	in full is Exercise 12.1.)

	Pointers: The Elem method returns the variable pointed to by a pointer, again as a
	reflect.Value. This operation would be safe even if the pointer value is nil, in which case
	the result would have kind Invalid, but we use IsNil to detect nil pointers explicitly so we
	can print a more appropriate message . We prefix the path with a "*" and parenthesize it to
	avoid ambiguity.

	Interfaces: Again, we use IsNil to test whether the interface is nil, and if not, we retrieve its
	dynamic value using v.Elem() and print its type and value.

*/

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

// display

// Display function
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			field := v.Field(i).String()
			search := v.Type().String()
			if strings.Contains(field,search) {
				fmt.Printf("%s = %v\n",path,formatAtom(v.Field(i)))
				return
			}
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			// Ejercicio 12.1
			if diff := reflect.ValueOf(key); diff.Kind() != reflect.String && diff.Kind() != reflect.Int &&
				diff.Kind() != reflect.Float64 && diff.Kind() != reflect.Bool && diff.Kind() != reflect.Float32 {
				display(fmt.Sprintf("%s", path), key)
			}
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

// Movie struct
type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

// Ejercicio 12.2
type cycle struct {
	value int
	tail *cycle
}

func main() {
	/*
	strangelove := Movie{
		Title:    "Dr.Strangelove",
		Subtitle: "How I learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Streangelove":           "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merking Muffley":      "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Rpper":   "Sterling Hayden",
			`Maj. T.J "King" Kong`:       "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin. )",
			"Best Picture (Nomin. )",
		},
	}

	Display("strangelove", strangelove)

	type testStruct struct{name,email string}
	testMap := make(map[testStruct]string)
	inStruct := testStruct{"First","email"}
	testMap[inStruct] = "first name"
	Display("mi prueba",testMap)
	*/
	var c cycle
	c = cycle{20,&c}
	Display("cycle struct",c)

}
