package main

import (
	"fmt"
	//"log"
)

/*
	Rece Conditions

	We can make a program concurrency-safe without making every con crete type in that
	program concurrency-safe. Indeed, concurrency-safe types are the exception rather than the
	rule, so you should access a variable concurrently only if the documentation for its type says
	that this is safe. We avoid concurrent access to most variables either by confining them to a
	single goroutine or by maintaining a higher-level invariant of mutual exclusion.

	A race condition is a situation in which the program does not give the correct result for some
	interleavings of the operations of multiple goroutines. Race conditions are pernicious because
	they may remain latent in a program and appear infrequently, perhaps only under heavy load
	or when using certain compilers, platforms, or architectures. This makes them hard to
	reproduce and diagnose.

	This program contains a particular kind of race condition called a data race. A data race
	occurs whenever two goroutines access the same variable concurrently and at least one of the
	accesses is a write.

	Things get even messier if the data race involves a variable of a type that is larger than a single
	machine word, such as an interface, a string , or a slice. This code updates x concurrently to
	two slices of dif ferent lengths:
	var x []int
	go func() { x = make([]int, 10) }()
	go func() { x = make([]int, 1000000) }()
	x[999999] = 1 // NOTE: undefined behavior; memory corruption possible!

	The value of x in the final statement is not defined; it could be nil, or a slice of length 10, or a
	slice of length 1,000,000. But recall that there are three parts to a slice: the pointer, the length,
	and the capacity. If the pointer comes from the first call to make and the length comes from
	the second, x would be a chimera, a slice whose nominal length is 1,000,000 but whose underlying array has only 10 elements.
	In this eventuality, storing to element 999,999 would clobber
	an arbitrary faraway memory location, with consequences that are impossible to predict and
	hard to debug and localize. This semantic minefield is called undefined behavior and is well
	known to C programmers; fortunately it is rarely as troublesome in Go as in C.

	The first way is not to write the variable. Consider the map below, which is lazily populated as
	each key is requested for the first time. If Icon is called sequentially, the program works fine,
	but if Icon is called concurrently, there is a data race accessing the map.

	The second way to avoid a data race is to avoid accessing the variable from multiple
	goroutines. This is the approach taken by many of the programs in the previous chapter. For
	example, the main goroutine in the concurrent web crawler (§8.6) is the sole goroutine that
	accesses the seen map, and the broadcaster goroutine in the chat server (§8.10) is the only
	goroutine that accesses the clients map. These variables are confined to a single goroutine.

*/

// bank1

var (
	deposits = make(chan int)
	balances = make(chan int)
	withmessage = make(chan W) // Parte del ejercicio 9.1
)

// Deposit function
func Deposit(amount int) { deposits <- amount }

// Balance function
func Balance() int { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case msg := <-withmessage: // Parte del ejericio 9.1
			ok := balance >= msg.amount
			if ok {
				balance -= msg.amount
			}
			msg.withdraw <- ok
		case balances <- balance:

		}
	}
}

func init() {
	go teller()
}

// Ejercicio 9.1

// W struct
type W struct {
	amount   int
	withdraw chan<- bool
}

// Withdraw function
func Withdraw(amount int) bool {
	sufficent := make(chan bool)
	withmessage <- W{amount,sufficent}
	return <-sufficent
}



func main() {
	/*
		done := make(chan struct{})
		//Accion 1
		go func(){
			Deposit(100)
			done <- struct{}{}
			fmt.Println("Action 1 done")
		}()
		// Accion 2
		go func(){
			Deposit(200)
			done <- struct{}{}
			fmt.Println("Action 2 done")
		}()
		// Accion 3
		go func(){
			Deposit(300)
			done <- struct{}{}
			fmt.Println("Action 3 done")
		}()
		// Esperando por liberar el canal
		<-done
		<-done
		<-done

		if got,want := Balance(),600; got < want {
			log.Println("Transition errors!")
		}
	*/

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
		fmt.Printf("Withdraw 100 OK? -> %t\n",Withdraw(200))
		done <- struct{}{}
	}()

	go func() {
		fmt.Printf("Withdraw 200 OK? -> %t\n",Withdraw(200))
		done <- struct{}{}
	}()

	<-done
	<-done
	<-done
	<-done

	fmt.Printf("\nFinal balance: %d\n", Balance())

}
