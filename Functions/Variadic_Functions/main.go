package main

/*
	A variadic function is one that can be called with varying numbers of arguments. The most
	familiar examples are fmt.Printf and its variants. Printf requires one fixed argument at the
	beginning, then accepts any number of subsequent arguments.

	To declare a variadic function, the type of the final parameter is preceded by an ellipsis, ‘‘...’’,
	which indicates that the function may be called with any number of arguments of this type

	Although the ...int parameter behaves like a slice within the function body, the type of a
	variadic function is distinct from the type of a function with an ordinary slice parameter.

	func f(...int) {}
	func g([]int) {}
	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"

	Variadic functions are often used for string formatting . The errorf function below constructs a formatted error message 
	with a line number at the beginning. The suffix f is a widely
	followed naming convention for variadic functions that accept a Printf-style format string .

	func errorf(linenum int, format string, args ...interface{}) {
		fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
		fmt.Fprintf(os.Stderr, format, args...)
		fmt.Fprintln(os.Stderr)
	}
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name) // "Line 12: undefined: count"
*/


import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"log"
	"os"
)


// Sum
func sum(vals...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}


// Ejercicio 5.15
func max(values...int) int {
	if len(values) <= 0 {
		log.Fatalln("At least 1 argument")
	}
	var maximun int = values[len(values) - 1]
	for _, value := range values {
		if value > maximun {
			maximun = value
		}
	}
	return maximun
}

func min(values...int) int {
	if len(values) <= 0 {
		log.Fatalln("At least 1 argument")
	}
	var minimun int = values[0]
	for _, value := range values {
		if value < minimun {
			minimun = value
		}
	}
	return minimun
}

// Ejercicio 5.16

func joint( step string,values...string) string {
	var result string
	for _,value := range values {
		 result += value + step
	}
	return result
}

// Ejercicio 5.17
func forEach(n *html.Node,f func(n *html.Node)) {
	if n != nil {
		f(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEach(c,f)
	}
	
}

// ElemnetsByTagName ...
func ElementsByTagName(doc *html.Node, name...string) []*html.Node {
	if len(name) == 0 {
		log.Fatalf("At least 1 name")
	}
	var found []*html.Node
	find := func(n *html.Node) {
			if n.Type == html.ElementNode {
				for _, tag := range name {
					if tag == n.Data {
						found = append(found,n)
					}
				}
			}
	}

	forEach(doc,find)
	return found
}

func main() {
	/*
	fmt.Printf("Result %d\n",sum(2))
	fmt.Printf("Result %d\n",sum(2,3,4))
	fmt.Printf("Result %d\n",sum(2,5,6,8))
	
	values := []int{1,2,3,5,6,7}
	fmt.Printf("Result %d\n",sum(values...))

	fmt.Printf("Maximun %d\n",max(2,5,6,7,0,1))
	fmt.Printf("Minimun %d\n",min(9,2,5,1,7,8))
	fmt.Printf("Maximun %d\n",max()) // sin argumentos
	
	fmt.Printf("Join => %s\n",joint("","H","o","l","a"," ","m","u","n","d","o"))
	*/
	url := os.Args[1:][0]
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Bad url")
	}
	doc, err := html.Parse(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Fatalf("Parse Error")
	}
	found := ElementsByTagName(doc,"a","img","h1","p")
	for _, node := range found {
		fmt.Printf("Founded %s\n",node.Data)
	}
}