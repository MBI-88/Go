package main

import (
	"os"
	"strings"
	"fmt"
	"io"
	"log"
	"net"
	"bufio"
	
)

// TableResponse struct
type TableResponse struct {
	name,address, output string
}

func (t *TableResponse) String() string {
	return fmt.Sprintf("%s | %s | %s",t.name,t.address,t.output)
}

func mustCopy(dst io.Reader, src *TableResponse) {
	buf := bufio.NewScanner(dst)
	for buf.Scan() {
		src.output = buf.Text()
		if err := buf.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n",src)
	}
	
}

func parse(hosts []string) []*TableResponse {
	resp := []*TableResponse{}
	for _,host := range hosts {
			source := strings.SplitN(host,"=",2)
			name,address := source[0],source[1]
			resp = append(resp,&TableResponse{name,address,""})
			

	}
	return resp
}

func main() {
	hosts := parse(os.Args[1:])
	fmt.Println("*----- Responses -----*")
	fmt.Println(" name  |  host |  time")
	for _, host := range hosts {
		conn, err := net.Dial("tcp",host.address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		mustCopy(conn,host)
		
		
	}
}