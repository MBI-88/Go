package main 

import (
	"path/filepath"
	"fmt"
	//"time"
	"sync"
	"os"
	"golang.org/x/net/html"
	"net/http"
	"log"
)

/*
 	Recall that after a channel has been closed and drained of all sent values, subsequent receive
	operations proceed immediately, yielding zero values. We can exploit this to create a broadcast 
	mechanism: don’t send a value on the channel, close it.

	We can add cancellation to the du program from the previous section with a few simple
	changes. First, we create a cancellation channel on which no values are ever sent, but whose
	closure indicates that it is time for the program to stop what it is doing . We also define a
	utility function, cancelled, that checks or polls the cancellation state at the instant it is called.

*/


var done = make(chan struct{})
var sema = make(chan struct{},20)

func cancelled() bool {
	select{
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage(nfiles,nbytes int64) {
	fmt.Printf("%d files %.1f GB\n",nfiles,float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir,entry.Name())
			go walkDir(subdir,n,fileSizes)
		}else {
			fileSizes <-entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func(){ <-sema }()
	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du:%v\n",err)
	}
	defer f.Close()
	entries, err := f.Readdir(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du:%v\n",err)
	}
	return entries
}

// Ejercicio 8.10

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

// Extract function
func Extract(url string, cancel <-chan struct{}) ([]string, error) {
	req, err := http.NewRequest("GET",url,nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		log.Fatalf("%v",err)
	}
	if resp == nil {
		return nil,nil
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := Extract(url,done)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}


func main() {
	/*
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	go func(){
		os.Stdin.Read(make([]byte,1))
		close(done)
	}()
	fileSize := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root,&n,fileSize)
	}
	go func(){
		n.Wait()
		close(fileSize)
	}()

	tick := time.Tick(500 * time.Millisecond)
	var nfiles,nbytes int64
	
loop:
	for {
		select{
		case <-done:
			for range fileSize {
				
			}
			return
		case size,ok := <-fileSize:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size 
		case <-tick :
			printDiskUsage(nfiles,nbytes)
		}
	}

	*/

	worklist := make(chan []string)
	var n int
	n++
	go func() { worklist <- os.Args[1:] }()
	go func() {
		os.Stdin.Read(make([]byte,1))
		close(done)
	}()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <- worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl2(link)
				}(link)

			}
		}
	}



}