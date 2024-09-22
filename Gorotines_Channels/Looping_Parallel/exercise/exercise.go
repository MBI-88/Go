package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// Reverb2 Ejercicio 8.4

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c *net.TCPConn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(c, input.Text(), 1*time.Second)
		}()
	}
	wg.Wait()
	c.CloseWrite()
	fmt.Println("Connection write close")
	
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8000")
	fmt.Println("[+] Server runing [+]")
	if err != nil {
		log.Fatal(err)
	}
	tcp,ok := listen.(*net.TCPListener)
	if !ok {log.Fatal(ok)}
	for {
		conn, err := tcp.AcceptTCP()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
