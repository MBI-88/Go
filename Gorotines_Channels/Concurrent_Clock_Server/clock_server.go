package main 

/*
	Netcat1 (for comunicating with Clock1)
	package main
	import (
		"io"
		"log"
		"net"
		"os"
	)
	func main() {
		conn, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		mustCopy(os.Stdout, conn)
	}

	func mustCopy(dst io.Writer, src io.Reader) {
		if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}

	The second client must wait until the first client is finished because the server is sequential; it
	deals with only one client at a time. Just one small change is needed to make the server concurrent: 
	adding the go keyword to the call to handleConn causes each call to run in its own
	goroutine.
*/

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"flag"
)

// Clock1

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c,time.Now().Format("15:04:05\n"))
		if err != nil {
			return 
		}
		time.Sleep(1 * time.Second)
	}
}

// Ejercicio 8.1 relacionado con clockwall

type portFlag struct {
	port string
}

func (p *portFlag) String() string {
	return p.port
}

// Set method
func (p *portFlag) Set(s string) error {
	var port string
	_,err := fmt.Sscanf(s,"%s",&port)
	if err != nil {return fmt.Errorf("Error %v",err)}
	p.port = port
	return nil
}


func getPort(name string, value string, uses string) *portFlag {
	pflag := portFlag{value}
	flag.CommandLine.Var(&pflag,name,uses)
	return &pflag
}	


// Ejercicio 8.2 (Hecho en XGO_Exercises)



func main() {
	
	port := getPort("port","8000","set port")
	flag.Parse()
	address := "localhost:"+port.String()
	listener,err := net.Listen("tcp",address)
	fmt.Printf("[+] Listener server %s\n",port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//handleConn(conn) // clock1
		go handleConn(conn) // clock2
	}
	

}