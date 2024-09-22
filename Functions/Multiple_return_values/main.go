package main

/*
	A function can return more than one result. We’ve seen many examples of functions from
	standard packages that return two values, the desired computational result and an error value
	or boolean that indicates whether the computation worked.

	In a function with named results, the operands of a return statement may be omitted. This is
	called a bare return.

	In functions like this one, with many return statements and several results, bare returns can
	reduce code duplication, but they rarely make code easier to understand. For instance, it’s not
	obvious at first glance that the two early returns are equivalent to return 0, 0, err (because
	the result variables words and images are initialized to their zero values) and that the final
	return is equivalent to return words, images, nil. For this reason, bare returns are best
	used sparingly.



*/

import (
	"os"
	"strings"
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	
)

// CountWordsAndImages return count words and images 
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s",err)
		return
	}
	words, images = countWordsAndImages(doc)
	return words, images, err
	
}

// Ejercicio 5.5

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return 
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style"){
		return countWordsAndImages(n.NextSibling)
	}else if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}else if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}
	cwords, cimages := countWordsAndImages(n.FirstChild)
	cswords, csimages := countWordsAndImages(n.NextSibling)

	words += cwords + cswords
	images += cimages + csimages
	return words,images

}


func main(){

	words, images, err := CountWordsAndImages(os.Args[1:][0])
	if err != nil {
		fmt.Printf("Value error: %d\n",err)
	}
	fmt.Printf("Total words: %d\nTotal images: %d\n",words,images)

}