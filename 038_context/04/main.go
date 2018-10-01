package main

import (
	"fmt"
	"time"
)

func main() {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

// gen is a broken generator that will leak a goroutine.
func gen() <-chan int {
	ch := make(chan int)
	go func() { // <- basically this go routine will spin on forever because we stop catching ints from the channel
		var n int // on line 11, and go to sleep, be this go routine will keep spinning and sending values through ch
		for {
			ch <- n
			n++
		}
	}()
	return ch
}
