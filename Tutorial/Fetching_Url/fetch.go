package main

import (
	"fmt"
	"io"
	"strings"

	//"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		// Eejercicio 2 agrear prefijo http:// a la entrada
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//b, err := ioutil.ReadAll(resp.Body)

		// Ejercicio 1 usar io.Copy(dst,src)
		b, err := io.Copy(os.Stdout, resp.Body)
		status := resp.Status
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
		fmt.Printf("%s\n %d\n", status, b)
	}
}
