package main

/*
	Observe the duplicated resp.Body.Close() call, which ensures that title closes the network connection on all execution paths, 
	including failures. As functions grow more complex and have to handle more errors, such duplication of clean-up logic may become 
	a maintenance problem. Let’s see how Go’s novel defer mechanism makes things simpler.

	Syntactically, a defer statement is an ordinary function or method call prefixed by the
	keyword defer. The function and argument expressions are evaluated when the statement is
	executed, but the actual call is deferred until the function that contains the defer statement
	has finished, whether normally, by executing a return statement or falling off the end, or
	abnormally, by panicking. Any number of calls may be deferred; they are executed in the reverse of the order 
	in which they were deferred.

	A defer statement is often used with paired operations like open and close, connect and disconnect, or lock and unlock 
	to ensure that resources are released in all cases, no matter how
	complex the control flow. The right place for a defer statement that releases a resource is
	immediately after the resource has been successfully acquired. In the title function below, a
	single deferred call replaces both previous calls to resp.Body.Close():

	The defer statement can also be used to pair ‘‘on entry’’ and ‘‘on exit’’ actions when debugging
	a complex function. The bigSlowOperation function below calls trace immediately, which
	does the ‘‘on ent ry’’ action then returns a function value that, when called, does the corresponding ‘‘on exit’’ action. 
	By deferring a call to the returned function in this way, we can
	instrument the entry point and all exit points of a function in a single statement and even pass
	values, like the start time, between the two actions. But don’t for get the final parentheses in
	the defer statement, or the ‘‘on entry’’ action will happen on exit and the on-exit action won’t
	happen at all!

	By naming its result variable and adding a defer statement, we can make the function print its
	arguments and results each time it is called.

	A deferred anonymous function can even change the values that the enclosing function
	returns to its caller.

	Because deferred functions aren’t executed until the very end of a function’s execution, a
	defer statement in a loop deserves extra scrutiny. The code below could run out of file
	descriptors since no file will be closed until all files have been processed:

	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close() // NOTE: risky; could run out of file descriptors
		// ...process f...
	}


*/


import (
	"io"
	"path"
	"log"
	"time"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"os"
)

// Title 1

 func forEachNode(node *html.Node, f func(node *html.Node), err error) {
	if node != nil {
		f(node)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c,f,err)
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html",url,ct)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url,err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
		n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc,visitNode, nil)
	return nil

}

// Title 2

func titleDefered(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url,ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
		n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}

	forEachNode(doc,visitNode, nil)
	return nil
}

// Trace

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s",msg)
	return func() {log.Printf("exit %s (%s)",msg, time.Since(start))}
}

func double(x int) (result int) {
	defer func() {fmt.Printf("double(%d) = %d\n",x,result)}()
	return x + x
}

func triple(x int) (result int) {
	defer func() {result += x}()
	return double(x)
}

// Fetch

func fetch(url string) (filename string,n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f,resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

// Ejercicio 5.18

func openFile(path string, resp *http.Response) (n int64, err error) {
	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	n, err = io.Copy(file, resp.Body)
	return n, nil

}

func newFetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0 , err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	n, err = openFile(local,resp)
	// defer func() { err = f.Close()}() 
	return local, n, err

}


func main() {
	/*
	url := os.Args[1:][0]
	title(url)
	
	bigSlowOperation()

	_ = double(3)

	fmt.Println(triple(4))

	url := os.Args[1:][0]
	fetch(url)
	*/

	url := os.Args[1:][0]
	path, n, er := newFetch(url)
	fmt.Printf("Path %s %d bytes %v\n",path,n,er)




}