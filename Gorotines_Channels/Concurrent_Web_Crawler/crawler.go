package main

/*
	We can limit parallelism using a buffered channel of capacity n to model a concurrency primitive called a counting semaphore.
	Conceptually, each of then vacant slots in the channel buffer
	represents a token entitling the holder to proceed. Sending a value into the channel acquires a
	token, and receiving a value from the channel releases a token, creating a new vacant slot.
	This ensures that at most n send scan occur without an intervening receive. (Although it
	might be more intuitive to treat filled slots in the channel buffer as tokens, using vacant slots
	avoids the need to fill the channel buffer after creating it.) Since the channel element type is
	not important, we’ll use struct{}, which has size zero.

	The program below shows an alternative solution to the problem of excessive concurrency.
	This version uses the original crawl function that has no counting semaphore , but calls it
	from one of 20 long-lived crawler goroutines, thus ensuring that at most 20 HTTP requests are
	active concurrently.

*/

import (
	"strings"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"golang.org/x/net/html"
)

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
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s as HTML: %v", url, err)
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

// crawl1

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// crawl2

var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

// Ejercicio 8.6

var (
	maxdepth int
	args     []string
)
var idchannel = make(chan struct{}, 2)

// Work struct
type Work struct {
	depht int
	url   string
}

func init() {
	flag.IntVar(&maxdepth, "depth", 3, "max depth to crawl")
	flag.Parse()
	args = flag.Args()
	
}

func limitCrawl(work Work) []Work {
	fmt.Printf("%s %d\n", work.url, work.depht)
	if work.depht >= maxdepth {
		return nil
	}
	idchannel <- struct{}{}
	list, err := Extract(work.url)
	<-idchannel
	if err != nil {
		log.Print(err)
	}
	works := []Work{}
	for _, link := range list {
		works = append(works, Work{work.depht + 1, link})
	}
	return works
}

// Ejercicio 8.7 (Custom)

const samedir = "./mirror"

func savaDoc(data []byte, name string ) {
	name = strings.Replace(name,"https://","",1)
	name = strings.Replace(name,".com","",1)
	name = strings.ReplaceAll(name,"/","")
	name = strings.ReplaceAll(name,"#","")
	dirname := samedir + "/" + name
	if _,err := os.Stat(dirname); os.IsNotExist(err) {
		os.Mkdir(dirname,os.ModePerm)
		file, er := os.Create(dirname + "/" + name + ".html")
		if er != nil {log.Fatal(er)}
		_,err := file.Write(data)
		if err != nil { log.Printf("Error in save file %s\n", err)}
		defer file.Close()
	}
}

func sameDomain(x, y string) bool {
	urlx, errx := url.Parse(x)
	urly, erry := url.Parse(y)

	if errx != nil && erry != nil {
		return false
	}

	if urlx.Host != urly.Host {
		return false
	}

	return true
}

func contructor(ws chan []Work) {
	links := <-ws
	domain := args[0]
	for _, link := range links {
		same := sameDomain(domain, link.url)
		if same {
			resp, err := http.Get(link.url)
			if err == nil && resp.StatusCode == 200 {
				data, _ := ioutil.ReadAll(resp.Body)
				savaDoc(data,link.url)
			}
			defer resp.Body.Close()
		}

	}
}

func main() {
	/*
		worklist := make(chan []string)
		go func() { worklist <- os.Args[1:] }()

		seen := make(map[string]bool)
		for list := range worklist {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
		}


		worklist := make(chan []string)
		var n int
		n++
		go func() { worklist <- os.Args[1:] }()
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



		// crawl3

		worklist := make(chan []string)
		unseenLinks := make(chan string)

		go func(){ worklist <- os.Args[1:]}()
		for i := 0; i < 20; i++ {
			go func() {
				for link := range unseenLinks {
					foundlinks := crawl(link)
					go func() {worklist <- foundlinks}()
				}
			}()
		}

		seen := make(map[string]bool)
		for list := range worklist {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		}


	*/
	if _, err := os.Stat(samedir); os.IsNotExist(err) {
		os.Mkdir(samedir,os.ModePerm)
	}
	workers := make(chan []Work)
	var n int
	n++
	go func() {
		works := []Work{}
		for _, elm := range args {
			works = append(works, Work{1, elm})
		}
		workers <- works
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-workers
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				n++
				go func(link Work) {
					workers <- limitCrawl(link)
				    contructor(workers)
				}(link)

			}
		}
	}
}
