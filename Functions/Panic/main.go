package main


/*
	During a typical panic, normal execution stops, all deferred function calls in that goroutine are
	executed, and the program crashes with a log message. This log message includes the panic
	value, which is usually an error message of some sort, and, for each goroutine, a stack trace
	showing the stack of function calls that were active at the time of the panic. This log message
	often has enough information to diagnose the root cause of the problem without running the
	program again, so it should always be included in a bug report about a panicking program.

	No t all panics come from the runtime. The bui lt-in panic function may be called directly; it
	accepts any value as an argument. A panic is often the best thing to do when some ‘‘impossible’’ 
	situation happens, for instance, execution reaches a case that logically can’t happen:
	
	switch s := suit(drawCard()); s {
		case "Spades": // ...
		case "Hearts": // ...
		case "Diamonds": // ...
		case "Clubs": // ...
		default:
			panic(fmt.Sprintf("invalid suit %q", s)) // Joker?
	}

	Consider the function regexp.Compile, which compiles a regular expression into an efficient
	form for matching. It returns an error if called with an ill-formed pattern, but checking this
	error is unnecessary and burdensome if the caller knows that a particular call cannot fail. In
	such cases, it’s reasonable for the caller to handle an error by panicking , since it is believed to
	be impossible

	func Compile(expr string) (*Regexp, error) { ...}
	func MustCompile(expr string) *Regexp {
		re, err := Compile(expr)
			if err != nil {
				panic(err)
			}
		return re
	}



*/



import (
	"fmt"
	"os"
	"runtime"
)

// Defer 1

func f(x int ) {
	fmt.Printf("f(%d)\n", x + 0/x)
	defer fmt.Printf("defer %d\n",x)
	f(x - 1)
}

// Defer 2

func printStack() {
	var buf [4096]byte 
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}



func main() {
	/*
		f(3)
		f(2)
		f(1)
	*/
	defer printStack()
	f(3)



}