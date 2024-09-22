package main

/*
	Functions are first-class values in Go: like other values, function values have types, and they
	may be assigned to variables or passed to or returned from functions. A function value may
	be called like any other function.
	func square(n int) int { return n * n }
	func negative(n int) int { return -n }
	func product(m, n int) int { return m * n }
	f := square
	fmt.Println(f(3)) // "9"
	f = negative
	fmt.Println(f(3)) // "-3"
	fmt.Printf("%T\n", f) // "func(int) int"
	f = product // compile error: can't assign f(int, int) int to f(int) int

	The zero value of a function type is nil. Calling a nil function value causes a panic:
	var f func(int) int
	f(3) // panic: call of nil function

	Function values may be compare d with nil:

	var f func(int) int
	if f != nil {
		f(3)
	}
	but they are not comparable, so they may not be compared against each other or used as keys
	in a map.

	The findLinks fun tion from Section 5.2 uses a helper function, visit, to visit all the nodes
	in an HTML document and apply an action to each one. Using a function value, we can separate the logic for tree traversal
	from the logic for the action to be applied to each node, letting
	us reuse the traversal with different actions.

	The functions also indent the output using another fmt.Printf trick. The * adverb in %*s
	prints a string padded with a variable number of spaces. The width and the string are
	provided by the arguments depth*2 and "".

*/

import (
	f "fmt"
	//"net/http"
	"os"
	"golang.org/x/net/html"
	"io/ioutil"
	"regexp"
)

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		f.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		f.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// Ejercicio 5.7

func openTag(n *html.Node) {
	switch n.Type {
		case html.ElementNode:
			if n.FirstChild != nil {
				f.Printf("%*s<%s ", depth * 2, "", n.Data)
				for _, attr := range n.Attr {
					f.Printf("%s=%q ",attr.Key,attr.Val)
				}
				f.Printf(">\n")
				depth++
			}else {
				f.Printf("%*s<%s ",depth * 2,"",n.Data)
				for _, attr := range n.Attr {
					f.Printf("%s=%q ",attr.Key,attr.Val)
				}
				f.Printf("/>\n")
			
			}

		case html.CommentNode:
			f.Printf("%*s <!--- %s\n",depth * 2,"",n.Data)
			depth++
			
		
		case html.TextNode:
			f.Printf("%*s %s\n",depth * 3,"",n.Data)
			depth++

	}

}

func closeTag(n *html.Node) {
	switch n.Type {
		case html.ElementNode:
			if n.FirstChild != nil {
				depth--
				f.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		
		case html.CommentNode:
			depth--
			f.Printf("%*s --->\n",depth * 2,"")
			

		case html.TextNode:
			depth--

	}
	

}

// Ejercicio 5.8

func findID(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}
	}
	return true
	
}

func makeHTML(n *html.Node, id string, findID func(n *html.Node, id string) bool) *html.Node {
	if findID != nil {
		if !findID(n, id) {
			return n
		}
		
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := makeHTML(c, id, findID)
		if node != nil {
			return node
		}
	}
	
	return nil
}

func getElementByID(doc *html.Node, id string) string {
	var result string
	if  node := makeHTML(doc, id, findID); node != nil {
		for _, attr := range node.Attr {
			if attr.Key == "id" {
				result = attr.Val
			}
		}
		return result
	}
	return "empty"
}

// Ejercicio 5.9

func extendString(s string) string {
	return s + s
}

func expand(s string, f func(string) string) string {
	re := regexp.MustCompile(`\$[^\s]+`)
	return re.ReplaceAllStringFunc(s, func(x string) string {
		return f(x[1:])
	})
}

func main() {
	/*
	values := os.Args[1:]
	response, err := http.Get(values[0])
	if err != nil {
		f.Errorf("Response %v\n", err)
		os.Exit(0)
	}

	doc, err := html.Parse(response.Body)
	response.Body.Close()
	if err != nil {
		f.Errorf("Error %v\n", err)
	}
	*/

	// forEachNode(doc, startElemnt, endElement)
	// forEachNode(doc, openTag, closeTag)
	// f.Printf("Node found: %s\n", getElementByID(doc,values[1])) 
	readstring, _ := ioutil.ReadAll(os.Stdin)
	f.Println(expand(string(readstring),extendString))	
}
