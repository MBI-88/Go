package main

import (
	"strings"
)

/*
	In Go, a map is a reference to a hash table, and a map type is written map[K]V, where K and V
	are the types of its keys and values. All of the keys in a given map are of the same type, and all
	of the values are of the same type, but the keys need not be of the same type as the values.

	The key type K must be comparable using ==, so that the map can test whether a given key is equal
	to one already within it.

	The built-in function make can be used to create a map:
	ages := make(map[string]int) // mapping from strings to ints

	We can also use a map literal to create a new map populated with some initial key/value pairs:

	ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	}

	This is equivalent to:

	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	so an alternative expression for a new empty map is map[string]int{}

	and removed with the built-in function delete:
	delete(ages, "alice") // remove element ages["alice"]

	All of these operations are safe even if the element isn’t in the map; a map lookup using a key
	that isn’t present returns the zero value for its type, so, for instance, the following works even
	when "bob" is not yet a key in the map because the value of ages["bob"] will be 0.

	ages["bob"] = ages["bob"] + 1 // happy birthday!

	But a map element is not a variable, and we cannot take its address:

	_ = &ages["bob"] // compile error: cannot take address of map element

	One reason that we can’t take the address of a map element is that growing a map might cause
	rehashing of existing elements into new storage locations, thus potentially invalidating the
	address.

	To enumerate all the key/value pairs in the map, we use a range-based for loop similar to
	those we saw for slices. Successive iterations of the loop cause the name and age variables to
	be set to the next key/value pair:

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	The zero value for a map type is nil, that is, a reference to no hash table at all

	var ages map[string]int
	fmt.Println(ages == nil) // "true"
	fmt.Println(len(ages) == 0) // "true"

	Most operations on maps, including lookup, delete, len, and range loops, are safe to perform on a nil map reference,
	since it behaves like an empty map. But storing to a nil map causes a panic:

	ages["carol"] = 21 // panic: assignment to entry in nil map

	You must allocate the map before you can store into it.

	For example, if the element type is numeric, you might have to distinguish
	between a non existent element and an element that happens to have
	the value zero, using a test like this:

	age, ok := ages["bob"]
	if !ok { // "bob" is not a key in this map; age == 0.}

	As with slices, maps cannot be compared to each other; the only legal comparison is wih nil.
	To test whether two maps contain the same keys and the same associated values, we must
	write a loop.

	func equal(x, y map[string]int) bool {
		if len(x) != len(y) {
			return false
		}
		for k, xv := range x {
			if yv, ok := y[k]; !ok || yv != xv {
				return false
			}
		}
		return true
	}

	The example below uses a map to record the number of times Add has been called with a given
	list of strings. It uses fmt.Sprintf to convert a slice of strings into a single string that is a
	suitable map key, quoting each slice element with %q to record string boundaries faithfully:

	var m = make(map[string]int)
	func k(list []string) string { return fmt.Sprintf("%q", list) }
	func Add(list []string) { m[k(list)]++ }
	func Count(list []string) int { return m[k(list)] }


*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
	//"sort"
)

// graph
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges

	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func charcount() {
	countletter := make(map[string]int)
	countdigit := make(map[int]int)
	countsymbols := make(map[rune]int)
	var invalid int
	var utflen [utf8.MaxRune + 1]int
	input := bufio.NewReader(os.Stdin)
	for {
		chart, nbyte, err := input.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if unicode.IsLetter(chart) {
			countletter[string(chart)]++
		}
		if unicode.IsDigit(chart) {
			countdigit[int(chart)]++
		}
		if chart == unicode.ReplacementChar && nbyte == 1 {
			invalid++
			continue
		}
		if unicode.IsSymbol(chart) {
			countsymbols[chart]++
		}
		utflen[nbyte]++

	}
	fmt.Printf("\nLetter\tcount\n")
	for c, n := range countletter {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("Digits\tcount\n")
	for c, n := range countdigit {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("Symbols\tcount\n")
	for c, n := range countsymbols {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

func wordfreq(path string) {
	file, err := os.Open(path)
	countwords := make(map[string]int)
	
	if err != nil {
		fmt.Fprintf(os.Stderr,"%v\n",err)
		os.Exit(1)
	}
	
	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanWords)
	for reader.Scan() {
		word := strings.ToLower(reader.Text())
		countwords[word]++
	}
	fmt.Print("Word freq\n")
	for word, count := range countwords {
		fmt.Printf("%q\t%d\n", word, count)
	}
}

func main() {
	/*
		// dedup
		var names []string
		ages := make(map[string]int)
		ages["alice"] = 31
		ages["charlie"] = 34

		for name := range ages {
			names = append(names,name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Printf("%s\t%d\n",name,ages[name])
		}

		seen := make(map[string]bool)
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			line := input.Text()
			if !seen[line] {
				seen[line] = true
				fmt.Println(line)
			}
		}
		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr,"dedup: %v\n",err)
			os.Exit(1)
		}

		// charcount

		counts := make(map[rune]int)
		var utflen [utf8.UTFMax + 1]int
		invalid := 0

		in := bufio.NewReader(os.Stdin)
		for {
			r,n,err := in.ReadRune()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr,"charcount: %v\n",err)
				os.Exit(1)
			}
			if r == unicode.ReplacementChar && n == 1 {
				invalid++
				continue
			}
			counts[r]++
			utflen[n]++
		}
		fmt.Printf("rune\tcount\n")
		for c, n := range counts {
			fmt.Printf("%q\t%d\n",c,n)
		}
		fmt.Print("\nlen\tcount\n")
		for i, n := range utflen {
			if i > 0 {
				fmt.Printf("%d\t%d\n",i,n)
			}
		}
		if invalid > 0 {
			fmt.Printf("\n%d invalid UTF-8 characters\n",invalid)
		}
	*/

	// Ejercicio 4.8
	// charcount()

	// Ejercicio 4.9
	wordfreq("wordfreq.txt")
}
