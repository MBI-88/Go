package main

import (
	"fmt"
)


/*
	You may wonder why the Balance method needs mutual exclusion, either channel-based or
	mutex-based. After all, unlike Deposit, it con sists only of a single operation, so there is no
	danger of another goroutine executing ‘‘in the middle’’ of it. There are two reasons we need a
	mutex. The first is that it’s equally important that Balance not execute in the middle of some
	other operation like Withdraw. The second (and more subtle) reason is that synchronization
	is about more than just the order of execution of multiple goroutines; synchronization also
	affects memory.



*/

// Prueba de efecto data race


func main(){
	var x, y int 
	sig := make(chan struct{})
	go func(){
		x = 1
		fmt.Print("y:",y," ")
		sig<- struct{}{}
	}()
	go func(){
		y = 1
		fmt.Print("x:",x," ")
		sig<- struct{}{}
	}()

	<-sig
	<-sig

}