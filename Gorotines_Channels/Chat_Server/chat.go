package main

import (
	"time"
	"strings"
	"bufio"
	"fmt"
	"log"
	"net"
)

/*
	There is one instance apiece of
	the main and broadcaster goroutines, and for each client connection there is one
	handleConn and one clientWriter goroutine. The broadcaster is a good illustration of how select
	is used, since it has to respond to three different kinds of messages.

	Next is the broadcaster. Its local variable clients records the current set of connected clients.
	The only information recorded about each client is the identity of its out going message channel.

	In addition, handleConn creates a clientWriter goroutine for each client that receives messages broadcast
	to the client’s out going message channel and writes them to the client’s network connection. The client writer’s
	loop terminates when the broadcaster closes the channel
	after receiving a leaving notification.

*/

// chat server

func chat() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	fmt.Println("[*] Server running on localhost:8000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " have arrived"
	entering <- ch
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- ch
	messages <- who + " have left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, "\t", msg)
	}
}

// Ejercicio 8.12 y 8.13

// Client Server
type Client struct {
	name string
	ch   chan<- string
}

var timer = 1 * time.Minute // 8.13

var (
	in   = make(chan Client)
	out   = make(chan Client)
	mess = make(chan string)
)

func broadcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case msg := <-mess:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <- in:
			var totalclient []string
			for client := range clients {
				totalclient = append(totalclient, client.name)
			}
			cli.ch <- fmt.Sprintf("Total %d %s",len(totalclient),strings.Join(totalclient,","))
			clients[cli] = true
		case cli := <- out:
			delete(clients,cli)
			close(cli.ch)
		}
	}
}

func handleClients(conn net.Conn) {
	signal := make(chan struct{})
	ch := make(chan string)
	ip := conn.RemoteAddr().String()
	go writerClient(conn,ch)
	ch <- "You are " + ip
	mess <- ip + " arrived!"
	in <- Client{ip,ch}
	input := bufio.NewScanner(conn)
	// Ejercicio 8.13
	go func(){
		for {
			if input.Scan(){
				mess <- ip + ": " + input.Text()
				signal <- struct{}{}
			}else {
				out <- Client{ip,ch}
				mess <- ip + " left!"
				conn.Close()
				return
			}
		}
	}()
	// Parte del ejercicio 8.12
	//for input.Scan(){
	//	mess <- ip + ": " + input.Text()
	//}
	//out <- Client{ip,ch}
	//mess <- ip + " left!"
	//conn.Close()

	for {
		select {
		case _,ok := <-signal:
			if !ok {
				conn.Close()
				return
			}
		case <-time.After(timer):
			conn.Close()
			return
			
		}
	}

}

func writerClient(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, "\t", msg)
	}

}

func chatServer() {
	lis,err := net.Listen("tcp","localhost:8000")
	if err != nil { log.Fatal(err)}
	fmt.Printf("Server running in %s ",lis.Addr())
	go broadcast()
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleClients(conn)
	}
}



func main() {
	//chat()
	chatServer()
}
