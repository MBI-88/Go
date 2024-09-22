package main

import (
	"fmt"
	"sync"
)

/*
	Since the Balance function only needs to read the state of the variable, it would in fact be safe
	for multiple Balance calls to run concurrently, so long as no Deposit or Withdraw call is running.
	In this scenario we need a special kind of lock that allows read-only operations to
	proceed in parallel with each other, but write operations to have fully exclusive access. This
	lock is called a multiple readers, single writer lock, and in Go it’s provided by sync.RWMutex:
	var mu sync.RWMutex
	var balance int
	func Balance() int {
		mu.RLock() // readers lock
		defer mu.RUnlock()
		return balance
	}

	RLock can be used only if there are no writes to shared variables in the critical section. In general,
	we should not assume that logically read-only functions or methods don’t also update
	some variables. For example, a method that appears to be a simple accessor might also increment an
	internal usage counter, or update a cache so that repeat calls are faster. If in doubt,
	use an exclusive Lock.

	It’s only profitable to use an RWMutex when most of the goroutines that acquire the lock are
	readers, and the lock is under contention, that is, goroutines routinely have to wait to acquire
	it. An RWMutex requires more complex internal bookkeeping , making it slower than a regular
	mutex for uncontended locks.


*/

var (
	mu      sync.RWMutex
	balance int
)

// Deposit function
func Deposit(amount int) int {
	mu.Lock()
	balance += amount
	defer mu.Unlock()
	return balance
}

// Balance function
func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func withdrawMessage(amount int) bool {
	mu.RLock()
	defer mu.RUnlock()
	if balance <= amount {
		return false
	}
	return true
}

func main() {
	done := make(chan struct{})
	go func() {
		fmt.Println("Deposit 100")
		Deposit(100)
		done <- struct{}{}
	}()
	go func() {
		fmt.Println("Deposit 200")
		Deposit(200)
		done <- struct{}{}
	}()
	go func() {
		fmt.Println("Deposit 300")
		Deposit(300)
		done <- struct{}{}
	}()
	go func() {
		fmt.Printf("Transition for 200 -> %t\n", withdrawMessage(200))
		done <- struct{}{}
	}()
	go func() {
		fmt.Printf("Transition for 300 -> %t\n", withdrawMessage(300))
		done <- struct{}{}
	}()

	<-done
	<-done
	<-done
	<-done
	<-done
	fmt.Printf("Total balance %d\n", Balance())
}
