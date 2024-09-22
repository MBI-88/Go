package main

import (
	"sync"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

/*
	The ioutil.ReadDir function returns a slice of os.FileInfo—the same information that a
	call to os.Stat returns for a single file. For each subdirectory, walkDir recursively calls itself,
	and for each file, walkDir sends a message on the fileSizes channel. The message is the size
	of the file in bytes.

	Since the program no longer uses a range loop, the first select case must explicitly test
	whether the fileSizes channel has been closed, using the two-result form of receive operation. 
	If the channel has been closed, the program breaks out of the loop. The labeled break
	statement breaks out of both the select and the for loop; an unlabeled break would break
	out of only the select, causing the loop to begin the next iteration.

*/

// du1

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d filess %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// du2

var verbose = flag.Bool("v", false, "show  verbose progress messages")

// du3

var sema = make(chan struct{},20)

func walkDir3(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents3(dir) {
		if entry.IsDir(){
			n.Add(1)
			subdir := filepath.Join(dir,entry.Name())
			go walkDir3(subdir,n,fileSizes)
		}else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents3(dir string) []os.FileInfo {
	sema <-struct{}{}
	defer func(){ <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3:%v\n", err)
		return nil
	}
	return entries
}



// Ejercicio 8.9

// RootDir struct
type RootDir struct {
	name string
	nfiles int64
	nbytes int64

}

func printRootDir(dirs []RootDir) {
	for _, dir := range dirs {
		fmt.Printf("%s %d %.1f GB\n",dir.name,dir.nfiles,float64(dir.nbytes)/1e9)
	}
}




func main() {
	/*
		flag.Parse()
		roots := flag.Args()
		if len(roots) == 0 {
			roots = []string{"."}
		}
		fileSizes := make(chan int64)
		go func() {
			for _, root := range roots {
				walkDir(root, fileSizes)
			}
			close(fileSizes)
		}()
		var nfiles, nbytes int64
		for size := range fileSizes {
			nfiles++
			nbytes += size
		}
		printDiskUsage(nfiles, nbytes)


		// du2

		flag.Parse()
		roots := flag.Args()
		if len(roots) == 0 {
			roots = []string{"."}
		}
		fileSizes := make(chan int64)
		go func() {
			for _, root := range roots {
				walkDir(root, fileSizes)
			}
			close(fileSizes)
		}()

		var tick <-chan time.Time
		if *verbose {
			tick = time.Tick(500 * time.Millisecond)
		}
		var nfiles, nbytes int64
	loop:
		for {
			select {
			case size, ok := <-fileSizes:
				if !ok {
					break loop
				}
				nfiles++
				nbytes += size
			case <-tick:
				printDiskUsage(nfiles, nbytes)
			}
		}


		// du3

		flag.Parse()
		roots := flag.Args()
		if len(roots) == 0 {
			roots = []string{"."}
		}
		fileSizes := make(chan int64)
		var n sync.WaitGroup
		for _, root := range roots {
			n.Add(1)
			go walkDir3(root,&n,fileSizes)
		}
		go func(){
			n.Wait()
			close(fileSizes)
		}()
		var nfiles,nbytes int64
		
		var tick <-chan time.Time
		if *verbose {
				tick = time.Tick(500 * time.Millisecond)
		}

		myloop:
		for {
			select{
			case size, ok := <-fileSizes:
				if !ok {
					break myloop
				}
				nfiles++
				nbytes += size
			case <-tick:
				printDiskUsage(nfiles,nbytes)
			}
		}


	*/

	flag.Parse()
	rootdir := flag.Args()
	
	var nroots []RootDir
	for _, dir := range rootdir {
		nroots = append(nroots,RootDir{dir,0,0})
	}
	if len(nroots) == 0 {
		nroots = []RootDir{RootDir{".",0,0}}
	}
	go func(){
		for {
			var n sync.WaitGroup
			for i := range nroots {
				n.Add(1)
				i := i
				go func(){
					fileSize := make(chan int64)
					var m sync.WaitGroup
					m.Add(1)
					go func(){
						walkDir3(nroots[i].name,&m,fileSize)
					}()
					go func(){
						m.Wait()
						close(fileSize)
						n.Done()
					}()
					var nfiles,nbytes int64
					for {
						size,ok := <-fileSize
						if !ok {
							break
						}
						nfiles++
						nbytes += size
					}
					nroots[i].nbytes = nbytes
					nroots[i].nfiles = nfiles
				}()
			}
			n.Wait()
			<-time.After(500 * time.Millisecond)
		}

	}()
	var keep int
	for keep < len(nroots) {
		<-time.After(500 * time.Millisecond)
		printRootDir(nroots)
		keep++
	}

}
