package main

import (
	"fmt"
	"sync"
)

/*
	However, the cost of enforcing mutually exclusive access to icons is that two goro utines
	cannot access the variable concurrently, even once the variable has been safely initialized and will
	never be modified again. This suggests a multiple-readers lock.

	The pattern above gives us greater concurrency but is complex and thus error-prone.
	Fortunately, the sync package provides a specialized solution to the problem of one-time initialization:
	sync.Once. Con ceptually, a Once consists of a mutex and a boolean variable that
	records whether initialization has taken place; the mutex guards both the boolean and the
	client’s data structures. The sole method, Do, accepts the initialization function as its argument.
	Let’s use Once to simplify the Icon function:
	var loadIconsOnce sync.Once
	var icons map[string]image.Image
	// Concurrency-safe.
	func Icon(name string) image.Image {
		loadIconsOnce.Do(loadIcons)
		return icons[name]
	}

	Each call to Do(loadIcons) locks the mutex and checks the boole an variable. In the first call,
	in which the variable is false, Do calls loadIcons and sets the variable to true. Subsequent
	calls do nothing, but the mutex synchronization ensures that the effects of loadIcons on
	memory (specifically, icons) become visible to all goroutines. Using sync.Once in this way,
	we can avoid sharing variables with other goroutines until they have been properly constructed

*/

// Ejercicio 9.2

var (
	pc [256]byte
	mu sync.Once
)

func initialization() {
	for _, i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount function
func PopCount(value uint64) int {
	mu.Do(initialization)
	var result int
	for i := 0; i < 8; i++ {
		result += int(pc[byte(value>>(i*8))])
	}
	return result
}

func main() {
	fmt.Printf("Result %d\n", PopCount(1<<8-1))
}
