package main

import "time"

/*
	NOTES :
	Sometimes a goroutine becomes blocked, we can avoid this by setting a
	timeout in the select
*/

func main() {
	c := make(chan int)
	o := make(chan bool)

	go func() {
		for {
			select {
			case v := <- c:
				println(v)
			case <- time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()

	<- o
}