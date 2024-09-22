package main

/*
	Named functions can be declared only at the package level, but we can use a function literal to
	denote a function value within any expression. A function literal is written like a function
	declaration, but without a name following the func keyword . It is an expression, and its value
	is called an anonymous function.

	Function literals let us define a function at its point of use. As an example, the earlier call to
	strings.Map can be rewritt en as
	strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")

	The squares example demonstrates that function values are not just code but can have state.
	The anonymous inner function can access and update the local variables of the enclosing
	function squares. These hidden variable references are why we classify functions as reference
	types and why function values are not comparable. Function values like these are implemented
	using a technique called closures, and Go programmers often use this term for function values.

	Wh en an anonymous function requires recursion, as in this example, we must first declare a
	variable, and then assign the anonymous function to that variable. Had these two steps been
	combined in the declaration, the function literal would not be within the scope of the variable
	visitAll so it would have no way to call itself recursively:
	visitAll := func(items []string) {
		// ...
		visitAll(m[item]) // compile error: undefined: visitAll
		// ...
	}

	5.6.1 Caveat: Capturing Iteration Variables

	Consider a program that must create a set of directories and later remove them. We can use a
	slice of function values to hold the clean-up operations. (For brevity, we have omitted all error
	handling in this example.)

	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d // NOTE: necessary!
		os.MkdirAll(dir, 0755) // creates parent directories too
		rmdirs = append(rmdirs, func() {
		os.RemoveAll(dir)
		})
	}
	// ...do some work...
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}

	You maybe wondering why we assigned the loop variabled to a new local variable dir within
	the loop body, instead of just naming the loop variable dir as in this subtly incorrect variant

	var rmdirs []func()
	for _, dir := range tempDirs() {
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
		os.RemoveAll(dir) // NOTE: incorrect!
		})
	}

	The reason is a consequence of the scope rules for loop variables. In the program immediately
	above , the for loop introduces a new lexical block in which the variable dir is declared. All
	function values created by this loop ‘‘capture’’ and share the same variable—an addressable
	storage location, not its value at that particular moment. The value of dir is updated in successive iterat ions,
	so by the time the cleanup functions are called, the dir variable has been
	updated several times by the now-completed for loop. Thus dir holds the value from the
	final iteration, and consequently all calls to os.RemoveAll will attempt to remove the same
	directory.
	
*/

import (
	"fmt"
	"log"
	"net/http"
	neturl "net/url"
	"io"
	"path/filepath"
	"strings"
	"os"
	"sort"
	"golang.org/x/net/html"
)

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// toposort

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete maths"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programing"},
	"formal languages":      {"discrete systems"},
	"networks":              {"operating systems"},
	"operating system":      {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seem := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seem[item] {
				seem[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

// links

func forEachNode(n *html.Node, f func(n *html.Node)) {
	if f != nil {
		f(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, f)
	}
}


// Extract links
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode)
	return links, nil
}

// Findlinks3

func breadthFirst(f func(item string) []string, worklist []string) {
	seem := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seem[item] {
				seem[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}


// Ejercicio 5.10
func topoSortmap(m map[string][]string) map[int]string {
	order := make(map[int]string)
	seem := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _,item := range items {
			if !seem[item] {
				seem[item] = true
				visitAll(m[item])
				order[len(order)] = item
				
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys,key)
	}
	visitAll(keys)
	return order
}

// Ejercicio 5.11
var prereqscyclic = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete maths"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programing"},
	"formal languages":      {"discrete systems"},
	"networks":              {"operating systems"},
	"operating system":      {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":   {"calculus"},
}

func topoSortExtended(m map[string][]string) ([]string,error) {
	var order []string
	seem := make(map[string]bool)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if !seem[item] {
				seem[item] = true
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				order = append(order,item)
			}else {
				cyclic := true
				for _, val := range order {
					if val == item {
						cyclic = false
					}
				}
				if cyclic {
					return fmt.Errorf("cyclic %s",item)
				}
			}
		}
		return nil
	}
	var keys []string 
	for key := range m {
		keys = append(keys,key)
	}
	sort.Strings(keys)
	err := visitAll(keys)
	if err != nil {
		return nil, err
	}
	return order,err
}

// Ejercicio 5.12


// Ejercicio 5.13
var outputdir string = "./out"

func sameDomain(x,y string) (bool,error) {
	urlx, errx := neturl.Parse(x)
	urly, erry := neturl.Parse(y)
	if errx != nil {
		return false, errx
	}
	if erry != nil {
		return false, erry
	}
	return urlx.Host == urly.Host, nil
}

func saveHTML(url string) error {
	path, err := neturl.Parse(url)
	if err != nil {
		return err
	}
	localdir := outputdir + path.Path
	
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if !strings.Contains(filepath.Base(localdir),".") {
		localdir += "/index.html"
	}

	errdir := os.MkdirAll(filepath.Dir(localdir),os.ModePerm)
	if errdir != nil {
		return errdir
	}

	file, errfile := os.Create(localdir)
	if errfile != nil {
		return errfile
	}

	_, errio := io.Copy(file,resp.Body)
	if errio != nil && file.Close() != nil {
		return errio
	}

	return nil
}

func newcrawl(url string) []string {
	list, err := Extract(url)
	if err != nil {
		log.Println(err)
	}
	for _, target := range list {
		sameHost, err := sameDomain(url,target)
		if err != nil {
			continue
		}else {
			if sameHost {
				saved := saveHTML(target)
				if saved != nil {
					log.Println(err)
				}
			}
		}
		
	}

	return list
}


// ***** Main function *****

func main() {
	/*
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	
	fmt.Printf("Output toposort\n %s\n",topoSort(prereqs))
	
	
	result, err := Extract(os.Args[1:][0])
	if err != nil {
		log.Fatalf("Error: %s\n",err)
	}
	fmt.Printf("***Links***\n %s\n",result)
	
	
	breadthFirst(crawl,os.Args[1:])
	
	
	fmt.Printf("Ouput toposortmap %v\n",topoSortmap(prereqs))
	
	result, err := topoSortExtended(prereqscyclic)
	if err != nil {
		log.Fatalf("Errors: %s",err)
	}
	for i,item := range result {
		fmt.Printf("%d\t:%s\n", 1 + i,item)
	}
	
	
	breadthFirst(newcrawl,os.Args[1:])
	*/
	
	// Ejercicio 5.14

	var keys []string 
	for key := range prereqs {
		keys = append(keys,key)
	}

	i := 0
	action := func(item string) []string {
		i++
		fmt.Printf("%d\t%s\n", i, item)
		return prereqs[item]
	}

	breadthFirst(action,keys)


}
