package main

/*
	A function for which failure is an expected behavior returns an addition al result, conventionally the last one. 
	If the failure has only one possible cause, the result is a boole an, usually
	called ok, as in this example of a cache lookup that always succeeds unless there was no entry
	for that key.
	
	value, ok := cache.Lookup(key)
	if !ok {
		// ...cache[key] does not exist...
	}

	The built-in type error is an interface type. We’ll see more of what this means and its implications for error handling in Chapter 7. 
	For now it’s enoug h to know that an error may be nil or non-nil, that nil implies success and non-nil implies failure , and that a non-nil 
	error has an error message string which we can obtain by calling its Error method or print by calling
	fmt.Println(err) or fmt.Printf("%v", err)

	By contrast, Go programs use ordinary control-flow mechanisms like if and return to
	respond to errors. This style undeniably demands that more attention be paid to error-handling logic, but that is precisely the point.


	Error-Handling Strategies

	The fmt.Errorf function formats an error message using fmt.Sprintf and returns a new
	error value. We use it to build descriptive errors by successively prefixing additional context
	information to the original error message.

	Because error messages are frequently chained together, message strings should not be capitalized and newlines should be avoided. 
	The resulting errors may be long, but the y will be selfcontained when found by tools like grep.

	A more convenient way to achieve the same effect is to call log.Fatalf. As with all the log
	functions, by default it prefixes the time and date to the error message.
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}	
	
	The default format is helpful in a long-running server, but less so for an interactive tool:
	2006/01/02 15:04:05 Site is down: no such domain: bad.gopl.io

	For a more attractive output, we can set the prefix used by the log package to the name of the
	command, and suppress the display of the date and time:
	log.SetPrefix("wait: ")
	log.SetFlags(0)

	Fo urth, in some cases, it’s sufficient just to log the error and then continue, perhaps with
	reduced functionality. Again there’s a choice between using the log package, which adds the
	usual prefix.
	if err := Ping(); err != nil {
		log.Printf("ping failed: %v; networking disabled", err)
	}	

	(All log functions append a newline if one is not already present.)

	End of File (EOF)

	The caller can detect this condition using a simple comparison, as in the loop below, which
	reads runes from the standard input. (The charcount program in Section 4.3 provides a more
	complete example.)


*/


import (
	"fmt"
	"os"
	"bufio"
	"io"
	//"errors"
)


func main(){
	//var EOF = errors.New("EOF")
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Errorf("read failed: %v", err)
			return
		}
		fmt.Println(r)
	}
	
}