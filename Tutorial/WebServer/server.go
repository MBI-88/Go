package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}

func showHeader(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", request.Method, request.URL, request.Proto)
	for k, v := range request.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", request.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", request.RemoteAddr)
	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range request.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	
}

// Lissajous makes a gif image and send info to the browser
func lissajous(out io.Writer, request *http.Request) {
	q := request.URL.Query().Get("cycles")
	tonumber,_ := strconv.Atoi(q)
	cycles := float64(tonumber)

	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5),
				blackIndex)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

/* Prueba de comentario en el interior de una funcion */
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Counter %d\n", count)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/show-header", showHeader)
	http.HandleFunc("/lissajous/", func(w http.ResponseWriter, request *http.Request) {
		lissajous(w, request)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
