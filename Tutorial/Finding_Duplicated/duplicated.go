package main

import (
	"bufio"
	"fmt"

	//"io/ioutil"
	"os"
	//"strings"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	// Primera version
	/*
		counts := make(map[string]int)
		input := bufio.NewScanner(os.Stdin)

		for input.Scan() {
			counts[input.Text()]++
		}

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	*/

	// Segunda version puede leer un archivo

	/*
		counts := make(map[string]int)
		files := os.Args[1:]
		if len(files) == 0 {
			countLines(os.Stderr,counts)

		}else {
			for _, arg := range files {
				f, err := os.Open(arg)
				if err != nil {
					fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
				}
				countLines(f,counts)
				f.Close()
			}
		}
		for line,n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n",n,line)
			}
		}
	*/

	// Version cargando en memoria
	/*
		counts := make(map[string]int)
		for _, filename := range os.Args[1:] {
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
				continue
			}
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	*/

	// Ejercicio

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stderr, counts)
	} else {
		for _, filename := range files {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
			} else {
				countLines(file, counts)
				file.Close()
				if len(counts) != 0 {
					fmt.Printf("%s\n",filename)
					for line, n := range counts {
						if n > 1 {
							fmt.Printf("%d\t%s\n", n, line)
						}
					}
				}
			}
		}

	}
	
	

}
