package main



/*
	A type assertion is an operation applied to an interface value. Syntactically, it looks like x.(T),
	where x is an expression of an interface type and T is a type, called the ‘‘asserted’’ type. A type
	assertion checks that the dynamic type of its operand matches the asserted type.

	Second, if instead the asserted type T is an interface type, then the type assertion checks
	whether x’s dynamic type satisfies T. If this check succeeds, the dynamic value is not extracted;
	the result is still an interface value with the same type and value components, but the result
	has the interface type T. In other words, a type assertion to an interface type changes the type
	of the expression, making a different (and usually larger) set of methods accessible, but it
	preserves the dynamic type and value components inside the interface value.

	Often we’re not sure of the dynamic type of an interface value, and we’d like to test whether it
	is some particular type. If the type assertion appears in an assignment in which two results are
	expected, such as the following declarations, the operation does not panic on failure but
	instead returns an additional second result, a boole an indicating success:

	var w io.Writer = os.Stdout
	f, ok := w.(*os.File) // success: ok, f == os.Stdout
	b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil

	Discriminating Errors with Type Assertions

	Consider the set of errors returned by file operations in the os package. I/O can fail for any
	number of reasons, but three kinds of failure often must be handled differently: file already
	exists (for create operations), file not found (for read operations), and per mission denied. The

	package os
	func IsExist(err error) bool
	func IsNotExist(err error) bool
	func IsPermission(err error) bool

	A more reliable approach is to represent structured error values using a dedicated type. The
	os package defines a type called PathError to describe failures involving an operation on a
	file path, like Open or Delete, and a variant called LinkError to describe failures of operations involving 
	two file paths, like Symlink and Rename. Here’s os.PathError.

	package os
	// PathError records an error and the operation and file path that caused it.
	type PathError struct {
		Op string
		Path string
		Err error
	}
	func (e *PathError) Error() string {
		return e.Op + " " + e.Path + ": " + e.Err.Error()
	}


	Querying Behaviors withh interface Type Assertions

	Although io.WriteString do cuments its assumption, few functions that call it are likely to
	do cument that they too make the same assumption. Defining a method of a particular type is
	taken as an implicit assent for a certain behavioral contract. Newcomers to Go, especially
	those from a background in strongly typed languages, may find this lack of explicitint ention
	unsettling, but it is rarely a problem in practice. With the exception of the empty interface
	interface{}, int erface types are seldom satisfied by unintended coincidence.


	Type Switches

	A switchst atement simplifies an if-else chain that performs a series of value equality tests.
	An analogous type switch statement simplifies an if-else chain of type assertions

	In its simplest form, a type switch looks like an ordinary switch statement in which the operand is x.(type)—that’s 
	literally the key word type—and each case has one or more types. A
	type switch enables a multi-way branch based on the interface value’s dynamic type. The nil
	case matches if x == nil, and the default case matches if no other case do es.

	Example Token-Based XML Decoding
*/

import (
	"fmt"
	"encoding/xml"
	//"io"
	//"os"
	"strings"
	"log"
)

func containsAll(x,y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

// Ejercicio 7.17 

func sliceToString(x []string,y []map[string]string) []string {
	newSlice := []string{}

	for i := range x {
		newSlice = append(newSlice,x[i])
		for k,v := range y[i] {
			newSlice = append(newSlice,k+"="+v)
		}
	}

	return newSlice
}

// Ejercicio 7.18

// Node interface
type Node interface{
	String() string

}

// CharData string type
type CharData string 

// String method
func (c CharData) String() string {
	return string(c)
}

// Element struct 
type Element struct{
	Type xml.Name
	Attr []xml.Attr
	Children []Node
}

// String method
func (e *Element) String() string {
	var attrs,children string
	for _, item := range e.Attr {
		attrs += fmt.Sprintf(" %s=%q",item.Name.Local,item.Value)
	}
	for _, child := range e.Children {
		children += child.String()
	}
	return fmt.Sprintf("<%s%s>%s</%s>",e.Type.Local,attrs,children,e.Type.Local)
}

// Parse xml
func Parse(dec *xml.Decoder) (Node, error) {
	var stack []*Element
	for {
		doc, err := dec.Token()
		if err != nil {
			return nil,err
		}
		switch tok := doc.(type){
		case xml.StartElement:
			ele := Element{
				Type:tok.Name,
				Attr:tok.Attr,
				Children:[]Node{},
			}
			if len(stack) > 0 {
				stack[len(stack) - 1].Children = append(stack[len(stack) - 1].Children,&ele)
			}
			stack = append(stack,&ele)
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("unespected tag closing")
			}else if len(stack) == 1 {
				return stack[0],nil
			}
			stack = stack[:len(stack) - 1]
		case xml.CharData:
			if len(stack) > 0 {
				stack[len(stack) - 1].Children = append(stack[len(stack) - 1].Children,CharData(tok))
			}
		}
	}
	
}

func main() {
	/*
	dec := xml.NewDecoder(os.Stdin)
	var stack []string 
	var attrs []map[string]string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}else if err != nil {
			fmt.Fprintf(os.Stderr,"xmlselect: %v\n",err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack,tok.Name.Local)
			attr := make(map[string]string)
			for _,at := range tok.Attr {
				attr[at.Name.Local] = at.Value
			}
			
			attrs = append(attrs,attr)
		case xml.EndElement:
			stack = stack[:len(stack) - 1]
		case xml.CharData:
			if containsAll(sliceToString(stack,attrs), os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack," "),tok)
			}
		
		}
	}
	*/
	tests := []struct {
		xml  string
		want string
	}{
		{
			"<xml></xml>",
			"<xml></xml>",
		},
		{
			"<html><body></body></html>",
			"<html><body></body></html>",
		},
		{
			"<html><body><div></div></body></html>",
			"<html><body><div></div></body></html>",
		},
		{
			`<html><body id="body"></body></html>`,
			`<html><body id="body"></body></html>`,
		},
		{
			`<html><body id="body"><div id="i" class="c"></div></body></html>`,
			`<html><body id="body"><div id="i" class="c"></div></body></html>`,
		},
	}
	for _, test := range tests {
		dec := xml.NewDecoder(strings.NewReader(test.xml))
		got, err := Parse(dec)
		if err != nil {
			log.Fatalf("Error %v\n",err)
		}
		if got.String() != test.want {
			log.Fatalf("Error parse")
		}
		fmt.Println(got)
	}


}