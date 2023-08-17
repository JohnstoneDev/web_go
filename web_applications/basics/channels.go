package main

/*
	NOTES :
	goroutines run in the same memory address space,
	so you have to maintain synchronization when you want to access shared memory.

	You have to use make tp create a new channel

	A channel is like a two way pipeline, used to send & receive data
	Channels use the <- operator to send & receive data
*/

/*
	Sending and receiving data in channels blocks by default, so it's much easier to use synchronous goroutines.
	What I mean by block is that a goroutine will not continue when receiving data from an empty channel,
	i.e (value := <-ch), until other goroutines send data to this channel. On the other hand, the goroutine will not continue until the data it sends to a channel, i.e (ch<-5), is received.
*/

import "fmt"

func sum(a []int, c chan int) {
	total := 0

	for _, v := range a {
		total += v
	}
	c <- total // send total to the channel
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}