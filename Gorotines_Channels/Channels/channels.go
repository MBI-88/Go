package main 

/*
	If goroutines are the activities of a concurrent Go program, channels are the connections
	between them. A channel is a communication mechanism that lets one goroutine send values
	to another goroutine. Each channel is a conduit for values of a particular type, called the
	channel’s element type. The type of a channel whose elements have type int is written
	chan int.

	As with maps, a channel is a reference to the data structure created by make. When we copy a
	channel or pass one as an argument to a function, we are copying a reference, so caller and
	callee refer to the same data structure. As with other reference types, the zero value of a channel is nil.

	Two channels of the same type may be compared using ==. The comparison is true if both are
	references to the same channel data structure. A channel may also be compared to nil.

	A channel has two principal operations, send and receive, collectively known as
	communications. A send statement transmits a value from one goroutine, through the channel, 
	to another goroutine executing a corresponding receive expression. Both operations are
	written using the <- op erator. In a send statement, the <- separates the channel and value 
	operands. In a receive expression, <- precedes the channel operand. A receive expression whose
	result is not used is a valid statement.

	To close a channel, we call the built-in close function:
	close(ch)

	A channel created with a simple call to make is called an unbuffered channel, but make accepts
	an optional second argument, an integer called the channel’s capacity. If the capacity is nonzero, 
	make creates a buffered channel.
	ch = make(chan int) // unbuffered channel
	ch = make(chan int, 0) // unbuffered channel
	ch = make(chan int, 3) // buffered channel with capacity 3


	Unbuffered Channels

	A send operation on an unbuffered channel blocks the sending goroutine until  another
	goroutine executes a corresponding receive on the same channel, at which point the value is
	transmitted and both goroutines may continue. Conversely, if the receive operation was
	attempted first, the receiving goroutine is blocked until another goroutine performs a send on
	the same channel.

	Messages sent over channels have two important aspects. Each message has a value, but
	sometimes the fact of communication and the moment at which it occurs are just as
	important. We call messages events when we wish to stress this aspect. When the event carries no additional 
	information, that is, its sole purpose is synchronization, we’ll emphasize this
	by using a channel whose element type is struct{}, though it’s common to use a channel of
	bool or int for the same purpose since done <- 1 is shorter than done <- struct{}{}


*/

import (
	"os"
	"log"
	"io"
	"net"
)

// Netcat3

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst,src); err != nil {
		log.Fatal(err)
		
	}
}


// Ejercicio 8.3


func main() {
	conn, err := net.Dial("tcp","localhost:8000")
	if err != nil {
		log.Fatal(err)

	}
	tcpconn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal(err)
	}
	
	done := make(chan struct{})
	go func(){
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	mustCopy(conn,os.Stdin)
	tcpconn.CloseWrite()
	<- done // wait for background goroutine to finish
	conn.Close()
	tcpconn.CloseRead()

}