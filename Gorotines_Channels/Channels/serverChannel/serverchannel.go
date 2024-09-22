package main 

import (
	"log"
	"fmt"
	"net"
	"bufio"
)

func handleConn(conn net.Conn) {
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		data := sc.Text()
		fmt.Println(data)
	}

}



func main() {
	l,err := net.Listen("tcp","localhost:8000")
	fmt.Println("[+] Listener on localhost:8000")
	defer l.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		data,err := l.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(data)
	}

}