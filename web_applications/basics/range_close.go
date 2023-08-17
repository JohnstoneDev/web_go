package main

/*
	NOTES :
	We can use range to operate on buffer channels like in slice & map
	It is impossible to send or receive data on a closed channel

	Another thing you need to remember is that channels are not like files.
	You don't have to close them frequently unless you are sure the channel is
	completely useless, or you want to exit range loops.
*/

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}

	// Remember to always close channels on producers & not consumers,
	// or it will go into panic mode

	close(c) // closes the channel
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	// will stop reading data from the channel until it's closed
	for i := range c {
		fmt.Println(i)
	}
}