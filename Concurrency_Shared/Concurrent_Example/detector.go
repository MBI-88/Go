package main 

import (
	"fmt"
	"log"
	"io/ioutil"
	"sync"
	"net/http"
	"time"
)



/*
	In this section, we’ll build a concurrent non-blocking cache, an abstraction that solves a
	problem that arises often in real-world concurrent programs but is not well addressed by
	existing libraries. This is the problem of memoizing a function, that is, caching the result of a
	function so that it need be computed only once. Our solution will be concurrency-safe and
	will avoid the contention associated with designs based on a single lock for the whole cache.

	This example shows that it’s possible to build many concurrent structures using either of the
	two approaches—shared variables and locks, or communicating sequential processes—
	without excessive complexity.

	It’s not always obvious which approach is preferable in a given situation, but it’s worth
	knowing how they correspond. Sometimes switching from one approach to the other can
	make your code simpler.

*/

// memo1 

// Memo1 struct
type Memo1 struct {
	f Func1
	cache map[string]result1
}

type result1 struct {
	value interface{}
	err error
}

// Func1 funct
type Func1 func(key string) (interface{},error)

// New1 function
func New1(f Func1) *Memo1 {
	return &Memo1{f:f, cache:make(map[string]result1)}
}

// Get method
func (memo *Memo1) Get(key string) (interface{},error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value,res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}


// memo2

// Memo2 struct
type Memo2 struct {
	f Func2 
	mu sync.Mutex
	cache map[string]result2
}

type result2 struct {
	value interface{}
	err error
}

// Func2 type function
type Func2 func(key string) (value interface{},err error)

// Get method
func (memo2 *Memo2) Get(key string) (value interface{},err error) {
	memo2.mu.Lock()
	res, ok := memo2.cache[key]
	if !ok {
		res.value, res.err = memo2.f(key)
		memo2.cache[key] = res
	}
	memo2.mu.Unlock()
	return res.value,res.err
}

// New2 function
func New2(f Func2) *Memo2 {
	return &Memo2{f:f, cache:make(map[string]result2)}
}

// memo3 

// Memo3 struct
type Memo3 struct {
	f Func3 
	mu sync.Mutex
	cache map[string]result3
}

type result3 struct {
	value interface{}
	err error
}

// Func3 type function
type Func3 func(key string) (value interface{},err error)

// Get method
func (memo3 *Memo3) Get(key string) (value interface{},err error) {
	memo3.mu.Lock()
	res, ok := memo3.cache[key]
	memo3.mu.Unlock()
	if !ok {
		res.value, res.err = memo3.f(key)
		memo3.mu.Lock()
		memo3.cache[key] = res
		memo3.mu.Unlock()
	}
	return res.value,res.err
}
// New3 function
func New3(f Func3) *Memo3 {
	return &Memo3{f:f, cache:make(map[string]result3)}
}

// memo4 

// Memo4 struct
type Memo4 struct {
	f Func4 
	mu sync.Mutex
	cache map[string]*entry
}

type result4 struct {
	value interface{}
	err error
}

type entry struct {
	res result4 
	ready chan struct{}
}

// Func4 type function
type Func4 func(key string) (value interface{},err error)

// Get method
func (memo4 *Memo4) Get(key string) (value interface{},err error) {
	memo4.mu.Lock()
	e := memo4.cache[key]
	if e == nil {
		e = &entry{ready:make(chan struct{})}
		memo4.cache[key] = e
		memo4.mu.Unlock()
		e.res.value, e.res.err = memo4.f(key)
		close(e.ready)
	}else {
		memo4.mu.Unlock()
		<-e.ready
	}
	return e.res.value,e.res.err
}

// New4 function
func New4(f Func4) *Memo4 {
	return &Memo4{f:f, cache:make(map[string]*entry)}
}


// memo5 

type request struct {
	key string 
	response chan<- result5
}

// Memo5 struct
type Memo5 struct {requests chan request}

type result5 struct {
	value interface{}
	err  error
}

type entry5 struct {
	res result5 
	ready chan struct{}
}

// New5 function
func New5(f Func5) *Memo5 {
	memo5 := &Memo5{requests:make(chan request)}
	go memo5.server(f)
	return memo5
}

// Func5 function
type Func5 func(key string) (value interface{},err error)

// Get method
func (memo5 *Memo5) Get(key string) (interface{},error) {
	response := make(chan result5)
	memo5.requests <- request{key,response}
	res := <-response 
	return res.value,res.err
}

// Close method
func (memo5 *Memo5) Close() {close(memo5.requests)}

func (memo5 *Memo5) server(f Func5) {
	cache := make(map[string]*entry5)
	for req := range memo5.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry5{ready:make(chan struct{})}
			cache[req.key] = e
			go e.call(f,req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry5) call(f Func5,key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry5) deliver(response chan<- result5) {
	<-e.ready
	response <-e.res
}


// Funcion a memorizar resultado

func httpGetBody(url string) (interface{},error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil,err 
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}



func main() {
   	men := New5(httpGetBody)
   for url := range incomingURLs() {
		start := time.Now()
		value, err := men.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url,time.Since(start),len(value.([]byte)))
   }

}