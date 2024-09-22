package main

/*
	The first goroutine, counter, generates the integers 0, 1, 2, ..., and sends them over a channel to
	the second goroutine, squarer, which receives each value, squares it, and sends the result over
	another channel to the third goroutine, printer, which receives the squared values and prints
	them. For clarity of this example, we have intentionally chosen very simple functions, though
	of course they are too computationally trivial to warrant their own goroutines in a realistic
	program.
	
	As you might expect, the program prints the infinite series of squares 0, 1, 4, 9, and so on.
	Pipelines like this may be found in long-running server programs where channels are used for
	lifelong communication between goroutines containing infinite loops. But what if we want to
	send only a finite number of values through the pipeline?

	If the sender knows that no further values will ever be sent on a channel, it is useful to communicate
	this fact to the receiver goroutines so that they can stop waiting . This is accomplished by closing the
	channel using the built-in close function: close(naturals)

	Unidirectional Channel Types

	This arrangement is typical. When a channel is supplied as a function parameter, it is nearly
	always with the intent that it be used exclusively for sending or exclusively for receiving.

	To document this intent and prevent misuse, the Go type system provides unidirectional channel types that expose
	only one or the other of the send and receive operations. The type
	chan<- int, a send-only channel of int, allows sends but not receives. Conversely, the type
	<-chan int, a receive-only channel of int, allows receives but not sends. (The position of the
	<- arrow relative to the chan keyword is a mnemonic.) Violations of this discipline are
	detected at compile time.

	The call counter(naturals) implicitly converts naturals, a value of type chan int, to the
	type of the parameter, chan<- int. The printer(squares) call does a similar implicit conversion to <-chan int. 
	Conversions from bidirectional to unidirectional channel types are
	permitted in any assignment. There is no going back, how ever: once you have a value of a
	unidirectional type such as chan<- int, there is no way to obtain from it a value of type
	chan int that refers to the same channel data structure.
	

	Buffered Channels
	
	A buffered channel has a queue of elements. The queue’s maximum size is determined when it
	is created, by the capacity argument to make. The statement below creates a buffered channel
	capable of holding three string values. Figure 8.2 is a graphical representation of ch and the
	channel to which it refers.
	ch = make(chan string, 3)

	A send operation on a buffered channel inserts an element at the back of the queue, and a
	receive operation removes an element from the front. If the channel is full, the send operation
	blocks its goroutine until space is made available by another goroutine’s receive . Conversely, if
	the channel is empty, a receive operation blocks until a value is sent by another goroutine
	
	In the unlikely event that a program needs to know the channel’s buffer capacity, it can be
	obtained by calling the built-in cap function: cap(ch) // "3"

	The choice between unbuffered and buffered channels, and the choice of a buffered channel’s
	capacity, may both affect the correctness of a program. Unbuffered channels give stronger
	synchronization guarantees because every send operation is synchronized with its corresponding receive; 
	with buffered channels, these operations are decoupled. Also, when we
	know an upper bound on the number of values that will be sent on a channel, it’s not unusual
	to create a buffered channel of that size and perform all the sends before the first value is
	received. Failure to allocate sufficient buffer capacity would cause the program to deadlock.

	

*/

import (
	"fmt"
)

// pipeline3

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x 
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}


func main() {
	/*
	// pipeline1
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func(){
		for x := 0;  ; x++ {
			naturals <- x
		}
	}()

	// squarer 
	go func(){
		for {
			x, ok := <- naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		fmt.Println(<- squares)
	}

	// pipeline2
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func(){
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	
	// squarer 
	go func(){
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}


	*/

	naturals := make(chan int)
	squares := make(chan int)
	
	go counter(naturals)
	go squarer(squares,naturals)
	printer(squares)


}