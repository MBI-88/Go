package main 

/*
	Notice that the third shout from the client is not dealt with until the second shout has petered
	out, which is not very realistic. A real echo would consist of the composition of the three independent 
	shouts. To simulate it, we’ll need more goroutines. Again, all we need to do is add
	the go keyword , this time to the call to echo.
*/

import (
	"strings"
	"fmt"
	"net"
	"time"
	"bufio"
	"log"
)

// reverb1 

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c,"\t",strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",shout)
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c,input.Text(),1 * time.Second)
	}
	c.Close()
}

func main() {
	listen, err := net.Listen("tcp","localhost:8000")
	fmt.Println("[+] Running server")
	if err != nil {
		log.Fatal(err)
	}
	
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		
		go handleConn(conn)


	}
	

}