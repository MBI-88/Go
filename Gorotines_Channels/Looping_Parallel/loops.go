package main

/*


 */

import (
	"sync"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Image ...
func Image(src image.Image) image.Image {
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect)
	} else {
		height = int(128 / aspect)
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

// ImageStream ..
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

// ImageFile2 ...
func ImageFile2(outfile, inflie string) (err error) {
	in, err := os.Open(inflie)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(outfile)
	if err != nil {
		return err
	}
	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", inflie, outfile, err)
	}
	return out.Close()
}

// ImageFile ...
func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile)
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}

// thumbnail1

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Print(err)
		}
	}
}

// thumbnail2

func makeThumbnails2(filename []string) {
	for _, f := range filename {
		go ImageFile(f)
	}
}

// thumbanil3

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) { // forma correcta de pasar argumentos a goroutines
			ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

// thumbnail4

func makeThumbnails4(filenames []string)  error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string){
			_, err := ImageFile(f)
			errors <- err 
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}

// thumbnail5 soluciona problema de la 4

func makeThumbnails5(filenames []string) (thumnailfile []string,err error) {
	type item struct {
		thumbfile string
		err error
	}
	ch := make(chan item,len(filenames))
	for _, f := range filenames {
		go func(f string){
			var it item 
			it.thumbfile, it.err = ImageFile(f)
			ch <- it 
		}(f)
		
	}
	for range filenames {
		it := <-ch 
		if it.err != nil {return nil,it.err}
		thumnailfile = append(thumnailfile,it.thumbfile)
	}
	return thumnailfile,nil
}

// thumnail6

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string){
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err  != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}
	go func(){
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
} 

func main() {
	arraynames := os.Args[1:]
	namethumnails, err := makeThumbnails5(arraynames)
	if err != nil {log.Fatal(err)}
	fmt.Println(namethumnails)
}
