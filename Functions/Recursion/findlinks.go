package main

/*
	Functions may be recursive, that is, they may call themselves, either directly or indirectly.
	Recursion is a powerful technique for many problems, and of course it’s essential for processing recursive data structures. 
	In Section 4.4, we used recursion over a tree to implement a simple insertion sort. In this section, we’ll use it again for processing HTML documents.



*/

import (
	"io"
	"strings"
	"os"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func fetch() io.ReadCloser {
	url := os.Args[1:][0]
	if !strings.HasPrefix(url,"http://"){
			url = "http://" + url
	}
	response, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}

	return response.Body
}


func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	/*
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links,c)
	}
	*/
	// Ejercicio 5.1
	links = visit(links, n.FirstChild)
	links = visit(links,n.NextSibling)
	
	return links
}

// Ejercicio 5.2
func population(data map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return data
	}
	if n.Type == html.ElementNode {
		data[n.Data]++
	}
	data = population(data,n.FirstChild)
	data = population(data,n.NextSibling)
	return data
}

// Ejercicio 5.3
func printContentHTML(data []string, n *html.Node) []string {
	if n == nil {
		return data
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		data = printContentHTML(data,n.NextSibling)
		return data
	}else if  n.Type == html.TextNode {
			data = append(data,n.Data)
	}
	data = printContentHTML(data,n.FirstChild)
	data = printContentHTML(data,n.NextSibling)
	return data
}

// Ejercicio 5.4
func extractsLink(data []string,n *html.Node) []string {
	if n == nil {
		return data
	}
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					data = append(data,attr.Val)
				}
			}
		case "img":
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					data = append(data,attr.Val)
				}
			}
		case "script":
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					data = append(data,attr.Val)
				}
			}
		case "source":
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					data = append(data,attr.Val)
				}
			}
		case "link":
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					data = append(data,attr.Val)
				}
			}
		case "form":
			for _, attr := range n.Attr {
				if attr.Key == "action" {
					data = append(data,attr.Val)
				}
			}
		case "html":
			for _, attr := range n.Attr {
				if attr.Key == "manifest" {
					data = append(data,attr.Val)
				}
			}
		case "video":
			for _, attr := range n.Attr {
				if attr.Key == "src" || attr.Key == "poster" {
					data = append(data,attr.Val)
				}
			}

		}
	}
	data = extractsLink(data,n.FirstChild)
	data = extractsLink(data,n.NextSibling)
	return data
}


// Imprime todo el arbol
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack,c)
	}
}


func main(){
	doc, err := html.Parse(fetch())
	if err != nil {
		fmt.Fprintf(os.Stderr,"findlinks1: %v\n", err)
		os.Exit(1)
	}

	//for _, link := range visit(nil,doc) {
	//	fmt.Println(link)
	//}

	//data := make(map[string]int)
	//for key, element := range population(data,doc) {
	//	fmt.Printf("%s %d\n",key,element)
	//}

	//for _, content := range printContentHTML(nil,doc){
	//	fmt.Println(content)
	//}
	
	for _, link := range extractsLink(nil,doc) {
		fmt.Println(link)
	}
	
}