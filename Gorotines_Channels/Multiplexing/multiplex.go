package main

/*
	The program below does the countdown for a rocket launch. The time.Tick function returns
	a channel on which it sends events periodically, acting like a metronome. The value of each
	event is a timestamp, but it is rarely as interesting as the fact of its delivery

	Now each iteration of the countdown loop needs to wait for an event to arrive on one of the
	two channels: the ticker channel if everything is fine (‘‘nominal’’ in NASA jargon) or an abort
	event if there was an ‘‘anomaly.’’ We can’t just receive from each channel because whichever
	operation we try first will block until completion. We need to multiplex these operations, and
	to do that, we need a select statement.

	select {
	case <-ch1:
		// ...
	case x := <-ch2:
		// ...use x...
	case ch3 <- y:
		// ...
	default:
		// ...
	}

	A select waits until a communication for some case is ready to proceed. It then performs
	that communication and executes the case’s associated statements; the other communications
	do not happen. A select with no cases, select{}, waits forever.

	Let’s return to our rocket launch program. The time.After function immediately returns a
	channel, and starts a new goroutine that sends a single value on that channel after the specified time.
	The select statement below waits until the first of two events arrives, either an abort
	event or the event indicating that 10 seconds have elapsed. If 10 seconds go by with no abort,
	the launch proceeds.

	If multiple cases are ready, select picks one at random, which ensures that every channel has
	an equal chance of being selected. Increasing the buffer size of the previous example makes its
	output non deterministic, because when the buffer is neither full nor empty, the select statement
	figuratively tosses a coin.

	The time.Tick function behaves as if it creates a goroutine that calls time.Sleep in a loop,
	sending an event each time it wakes up. When the countdown function above returns, it stops
	receiving events from tick, but the ticker goroutine is still there , trying in vain to send on a
	channel from which no goroutine is receiving—a goroutine leak (§8.4.4)

	The Tick function is convenient, but it’s appropriate only when the ticks will be needed
	throughout the lifetime of the application. Otherwise, we should use this pattern
	ticker := time.NewTicker(1 * time.Second)
	<-ticker.C // receive from the ticker's channel
	ticker.Stop() // cause the ticker's goroutine to terminate

	Sometimes we want to try to send or receive on a channel but avoid blocking if the channel is
	not ready—a non-blocking communication. A select statement can do that too. A select
	may have a default, which specifies what to do when none of the other communications can
	proceed immediately.
*/

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// countdown 1
func launch() {
	fmt.Println("Lift off!")
}

func countDown1() {
	fmt.Println("Commencing countdown")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

// countdown 2

func countDown2() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abort")
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Countdown is working")
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

// countdown 3

func countDown3() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abort")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			fmt.Println("tick is sending")
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}
	launch()
}

// Ejercicio 8.8

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func echoServer() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[+] Server is listening")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go func(c net.Conn) {
			signal := make(chan struct{})
			input := bufio.NewScanner(c)

			go func() {
				for {
					if input.Scan() {
						signal <- struct{}{}
					} else {
						c.Close()
						return
					}
				}
			}()

			for {
				select {
				case _, ok := <-signal:
					if !ok {
						c.Close()
						return
					}
					echo(c, input.Text(), 1*time.Second)
				case <-time.After(10 * time.Second):
					fmt.Fprintln(c, "[*] Connection time out!")
					c.Close()
				}
			}
		}(conn)

	}
}

func main() {
	// countDown1()
	// countDown2()
	// countDown3()
	echoServer()
}
