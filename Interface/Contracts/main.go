package main

/*
	Interface types express generalizations or abstractions about the behaviors of other types. By
	generalizing, interfaces let us write functions that are more flexible and adaptable because they
	are not tied to the details of one particular implementation.

	There is another kind of type in Go called an interface type. An interface is an abstract type. It
	doesn’t expose the representation or internal structure of its values, or the set of basic 
	operations they support; it reveals only some of their methods. When you have a value of an
	interface type, you know nothing about what it is; you know only what it can do, or more
	precisely, what behaviors are provided by its methods.


*/


import (
	"strings"
	//"io/ioutil"
	"io"
	"bytes"
	"bufio"
	"fmt"
)

// bycounter

// ByteCounter type int
type ByteCounter int 

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// Ejercicio 7.1

// CountWords type int
type CountWords int

// CountLines type int
type CountLines int

// Write words
func (s *CountWords) Write(p []byte) (int,error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	i := 0
	for scanner.Scan() {
		i++
	}
	*s += CountWords(i)
	return i, nil
}

func (s *CountLines) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		i++
	}
	*s += CountLines(i)
	return i, nil
}

// Ejercicio 7.2

// CounterByte ...
type CounterByte struct {
	write io.Writer
	count int64
}

func (c *CounterByte) Write(p []byte) (int, error) {
	c.write.Write(p)
	c.count += int64(len(p))
	return len(p), nil
}

// CountingWriter function
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	counter := CounterByte{
		write: w, 
		count: 0,
	}
	
	return &counter,&counter.count
}

// Ejercicio 7.3

type tree struct {
	value int 
	left, right *tree
}


func (t *tree) String() (result string) {
	values := appendValues([]int{},t)
	strs := []string{}
	for _, v := range values {
		strs = append(strs,fmt.Sprint(v))
	}
	return strings.Join(strs," ")
}

// Sort ...
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root  = add(root, v)
	}
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)

	}else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values,t.value)
		values = appendValues(values, t.right)
	}
	return values
}




func main() {
	/*
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0 
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	var s CountWords
	wcount, _ := s.Write([]byte("hola algo mas"))
	fmt.Println(wcount)

	var l CountLines
	lcount,_ := l.Write([]byte("hola mundo mas\nSegunda linea"))
	fmt.Println(lcount)

	w, count := CountingWriter(ioutil.Discard)
	fmt.Fprint(w,"hello")
	fmt.Println(*count)

	*/
	var t *tree 
	t = add(t, 1)
	t = add(t, 2)
	t = add(t, 3)
	fmt.Println(t.String()) 
	

	
	



}