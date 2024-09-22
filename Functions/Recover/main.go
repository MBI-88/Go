package main

/*
	If the bui lt-in recover function is called within a deferred function and the function containing the defer statement is panicking , 
	recover ends the current state of panic and returns the
	panic value. The function that was panicking does not continue where it left off but returns
	normally. If recover is called at any other time, it has no effect and returns nil.

	func Parse(input string) (s *Syntax, err error) {
		defer func() {
			if p := recover(); p != nil {
				err = fmt.Errorf("internal error: %v", p)
			}
		}()
	// ...parser...
	}

	Recovering from a panic within the same package can help simplify the handling of complex
	or unexpected errors, but as a general rule, you should not attempt to recover from another
	package’s panic. Public APIs should report failures as errors. Similarly, you should not
	recover from a panic that may pass through a function you do not maintain, such as a caller-provided callback, 
	since you cannot reason about its safety.

	For all the above reasons, it’s safest to recover selectively if at all. In other words, recover only
	from panics that were intended to be recovered from, which should be rare. This intention
	can be encoded by using a distinct, unexported type for the panic value and testing whether
	the value returned by recover has that type. If so, we report the panic as an ordinary error; 
	if not, we call panic with the same value to resume the state of panic

*/


import (
	"log"
	//"net/http"
	//"os"
	"fmt"
	"golang.org/x/net/html"
)

// Title 3

func forEachNode(node *html.Node, f func(node *html.Node)) {
	if node != nil {
		f(node)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c,f)
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func(){
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()
	forEachNode(doc,func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
		n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	})
	if title == "" {
		return "",fmt.Errorf("no title element")
	}
	return title, nil

}

// Ejercicio 5.19

func findNull(v string, f func(v string)) {
	if v != "" {
		f(v)
	}
	fmt.Printf("Value of string %s\n",v)
}

func panicrecover(value string) {
	defer func() {
		if p := recover(); p != nil {	
			log.Fatalf("Recover errors %v",p)
		}
	}()
	
	findNull(value, func(value string){
		if value == "call panic"{
			panic("Panic was called")
		}else {
			fmt.Println(value)
		}
	})

}




func main() {
	/*
	url  := os.Args[1:][0]
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("Url error %s\n", err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("Html error %s\n",err)
	}

	soleTitle(doc)
	*/
	panicrecover("call panic")


}