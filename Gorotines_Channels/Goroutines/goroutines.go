package main

/*
	When a program starts, its only goroutine is the one that calls the main function, so we call it
	the main goroutine. New goroutines are created by the go statement. Syntactically, a go statement 
	is an ordinary function or method call prefixed by the keyword go. A go statement
	causes the function to be called in a newly created goroutine. The go statement it self completes immediately:
	f() // call f(); wait for it to return
	go f() // create a new goroutine that calls f(); don't wait

	The main function then returns. When this happens, all goroutines are abruptly terminated
	and the program exits. Other than by returning from main or exiting the program, there is no
	programmatic way for one goroutine to stop another, but as we will see later, there are ways to
	communicate with a goroutine to request that it stop itsself.

*/


import (
	"time"
	"fmt"
)

// Spinner

func spinner(delay time.Duration) {
	for {
		for _,r := range `-\|/` {
			fmt.Printf("\r%c",r)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}

func main(){
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n",n,fibN)

}