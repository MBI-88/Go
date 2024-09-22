package main

/*
	A string is an immutable sequence of bytes. Strings may contain arbitrary data, including
	bytes with value 0, but usually they contain human-readable text. Text strings are conventionally
	interpreted as UTF-8-encoded sequences of Unicode code points (runes)

	Because Go source files are always encoded in UTF-8 and Go text strings are conventionally
	interpreted as UTF-8, we can include Unicode code points in string literals.

	\a ‘‘alert’’ or bell
	\b backspace
	\f form feed
	\n newline
	\r carriage return
	\t tab
	\v vertical tab
	\' single quote (only in the rune literal '\'')
	\" double quote (only within "..." literals)
	\\ backslash

	Arbitrary bytes can also be included in literal strings using hexadecimal or octal escapes. A
	hexadecimal escape is written \xhh, with exactly two hexadecimal digits h (in upper or lower
	case). An octal escape is written \ooo with exactly three octal digits o (0 through 7) not
	exceeding \377. Both denote a single byte with the specified value. Later, we’ll see how to
	enco de Unico deco de points numerically in string literals

	UTF-8

	0xxxxxx runes 0−127 (ASCII)
	11xxxxx 10xxxxxx 128−2047 (values <128 unused)
	110xxxx 10xxxxxx 10xxxxxx 2048−65535 (values <2048 unused)
	1110xxx 10xxxxxx 10xxxxxx 10xxxxxx 65536−0x10ffff (other values unused)

	Strings and Byte Slices

	Four standard packages are particularly important for manipulating strings: bytes, strings,
	strconv, and unicode. The strings package provides many functions for searching, replacing, comparing ,
	trimming, splitting, and joining strings.

	The bytes package has similar functions for manipulating slices of bytes, of type []byte,
	which share some properties with strings. Because strings are immutable, building up strings
	incrementally can involve a lot of allocation and copying. In such cases, it’s more efficient to
	use the bytes.Buffer type.

	The strconv package provides functions for converting boolean, integer, and floating-point
	values to and from their string representations, and functions for quoting and unquoting
	strings.

	The unicode package provides functions like IsDigit, IsLetter, IsUpper, and IsLower for
	classifying runes. Each function takes a single rune argument and returns a boolean. Conversion functions
	like ToUpper and ToLower convert a rune into the given case if it is a letter. All
	these functions use the Unico de standard categories for letters, digits, and so on

	The path and path/filepath packages provide a more general set of functions for manipulating hierarchical names.
	The path package works with slash-delimited paths on any platform. It shouldn’t be used for file names, but it is
	appropriate for other domains, like the path component of a URL.
	By contrast, path/filepath manipulates file names using the rules for
	the host platform, such as /foo/bar for POSIX or c:\foo\bar on Micros oft Windows

	A string contains an array of bytes that, once created, is immutable. By contrast, the elements
	of a byte slice can be freely modified

	Conceptually, the []byte(s) conversion allocates a new byte array holding a copy of the bytes
	of s, and yields a slice that references the entirety of that array. An optimizing compiler maybe
	able to avoid the allocation and copying in some cases, but in general copying is required to
	ensure that the bytes of s remain unchanged even if those of b are subsequently modified. The
	conversion from byte slice back to string with string(b) also makes a copy, to ensure
	immutability of the resulting string s2.

	In addition to conversions between strings, runes, and bytes, it’s often necessary to convert
	between numeric values and their string representations. This is done with functions from the
	strconv package.
	To convert an integer to a string , one option is to use fmt.Sprintf; another is to use the function strconv.Itoa (‘‘integer to ASCII’’):
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"

	FormatInt and FormatUint can be used to format numbers in a different base:
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"

	The fmt.Printf verbs %b, %d, %u, and %x are often more convenient than Format functions,
	especially if we want to include additional information besides the number:
	s := fmt.Sprintf("x=%b", x) // "x=1111011"

	To parse a string representing an integer, use the strconv functions Atoi or ParseInt, or
	ParseUint for unsigned integers:
	x, err := strconv.Atoi("123") // x is an int
	y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits

	The third argument of ParseInt gives the size of the integer type that the result must fit into;
	for example, 16 implies int16, and the special value of 0 implies int. In any case, the type of
	the result y is always int64, which you can then convert to a smaller type.
*/

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

// Usando el modulo strings
func basenameStrconv(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// Ejercicio 3.10
func commaBuffer(s string) string {
	var buffer bytes.Buffer
	i := (3 - utf8.RuneCountInString(s)%3) % 3
	for _, r := range s {
		if i == 3 {
			buffer.WriteByte(',')
			i = 0
		}
		buffer.WriteRune(r)
		i++

	}
	return buffer.String()
}

// Ejercicio 3.11
func commaFloat(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	if point := strings.Index(s, "."); point != -1 {
		if substring := len(s[:point]); substring > 3 {
			for i, r := range s[:point] {
				if i%3 == 0 && i != 0 {
					buf.WriteByte(',')
					buf.WriteRune(r)
				} else {
					buf.WriteRune(r)
				}
			}
		} else {
			for _, i := range s[:point] {
				buf.WriteRune(i)
			}

		}
		if substring := len(s[point:]); substring > 3 {
			for i, r := range s[point:] {
				if i%3 == 0 && i != 0 {
					buf.WriteRune(r)
					buf.WriteByte(',')
				} else {
					buf.WriteRune(r)
				}
			}
		} else {
			for _, i := range s[point:] {
				buf.WriteRune(i)
			}
		}
	} else {
		i := (3 - utf8.RuneCountInString(s)%3) % 3
		for _, r := range s {
			if i == 3 {
				buf.WriteByte(',')
				i = 0
			}
			buf.WriteRune(r)
			i++
		}
	}
	return buf.String()
}

// Ejercicio 3.12
func appearCount(valor string) map[rune]int {
	valor = strings.ToLower(valor)
	array := make(map[rune]int)
	for _, r := range valor {
		array[r]++
	}
	return array
}

func anagramas(stringA, stringB string) bool {
	arrayA := appearCount(stringA)
	arrayB := appearCount(stringB)
	if len(arrayA) != len(arrayB) {
		return false
	}
	for _, r := range stringA {
		if arrayA[r] != arrayB[r] {
			return false
		}
	}

	return true

}

func main() {
	/*
		s := "hello, world"
		fmt.Println(len(s))
		fmt.Println(s[0],s[7])
		fmt.Println(s[0:5])
		fmt.Println("good " + s[5:])

		s := "abc"
		b := []byte(s)
		s2 := string(b)
		fmt.Printf("string => %s byte => %b stirng => %s\n",s,b,s2)
		fmt.Printf("%s\n",basename("hola mundo / Go run ."))
		fmt.Printf("%s\n",basenameStrconv("Hola / mundo Go"))
		fmt.Printf("%s\n",comma("1234895"))
	*/
	//fmt.Printf("%s\n",commaBuffer("12348"))
	//fmt.Printf("%s\n",commaFloat("1234567126.4567"))
	fmt.Printf("Anagrams %s <--> %s %t\n", "Nepal", "panel", anagramas("Nepal", "panel"))

}
