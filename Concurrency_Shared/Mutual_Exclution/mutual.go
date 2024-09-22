package main

import (
	"sync"
)

/*
	This pattern of mutual exclusion is so useful that it is supported directly by the Mutex type
	from the sync package. Its Lock method acquires the token (called a lock) and its Unlock
	method releases it: (bank 2)

	Each time a goroutine accesses the variables of the bank (just balance here), it must call the
	mutex’s Lock method to acquire an exclusive lock. If some other goroutine has acquired the
	lock, this operation will block until the other goroutine calls Unlock and the lock becomes
	avai lable again. The mutex guards the shared variables. By convention, the variables guarded
	by a mutex are declared immediately after the declaration of the mutex itself. If you deviate
	from this, be sure to document it.

	The region of code between Lock and Unlock in which a goroutine is free to read and modify
	the shared variables is called a critical section. The lock holder’s call to Unlock happens before
	any other goroutine can acquire the lock for itself. It is essential that the goroutine release the
	lock once it is finished, on all paths through the function, including error paths.

	The bank program above exemplifies a common concurrency pattern. A set of exported functions encapsulates
	one or more variables so that the only way to access the variables is through
	these functions (or methods, for the variables of an object). Each function acquires a mutex
	lock at the beginning and releases it at the end, there by ensuring that the shared variables are
	not accessed concurrently. This arrangement of functions, mutex lock, and variables is called
	a monitor. (This older use of the word ‘‘monitor’’ inspired the term ‘‘monitor goroutine.’’ Both
	uses share the meaning of a broker that ensures variables are accessed sequentially.)

	Since the critical sections in the Deposit and Balance functions are so short—a single line, no
	branching—calling Unlock at the end is straightforward. In more complex critical sections,
	especially those in which errors must be dealt with by returning early, it can be hard to tell that
	calls to Lock and Unlock are strictly paired on all paths. Go’s defer statement comes to the
	rescue: by deferring a call to Unlock, the critical section implicitly extends to the end of the
	current function, freeing us from having to remember to insert Unlock calls in one or more
	places far from the call to Lock.
	func Balance() int {
		mu.Lock()
		defer mu.Unlock()
		return balance
	}
	In the example above , the Unlock executes after the return statement has read the value of
	balance, so the Balance function is concurrency-safe. As a bonus, we no longer need the
	local variable b.

	There is a good reason Go’s mutexes are not re-entrant. The purpose of a mutex is to ensure
	that certain invariants of the shared variables are maintained at critical points during program
	execution. One of the invariants is ‘‘no goroutine is accessing the shared variables,’’ but there
	may be additional invariants specific to the data structures that the mutex guards. When a
	goroutine acquires a mutex lock, it may assume that the invariants hold. While it holds the
	lock, it may update the shared variables so that the invariants are temporarily violated.
	However, when it releases the lock, it must guarantee that order has been restored and the
	invariants hold once again. Alt hough a re-entrant mutex would ensure that no other
	goro utines are accessing the shared variables, it cannot protect the additional invariants of
	those variables.

*/


// bank 2

var (
	sema    = make(chan struct{}, 1)
	balance int
)

// Deposit function
func Deposit(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

// Balance fucntion
func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

// bank 3

var mu sync.Mutex

// DepositMux function
func DepositMux(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}
// BalanceMux function
func BalanceMux() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func main() {
}
