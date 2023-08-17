package main

/*
	NOTES :
	Buffered channels can store more than a singl element i.e : ch := make(chan bool 4)
	will create a channel that can store 4 boolean elements

	So, we can send multiple elements into the channel without blocking, but the goroutine
	will be blocked when you try to send a fifth element & no goroutine receives it
*/

import "fmt"

func main() {
    c := make(chan int, 2) // change 2 to 1 will have runtime error, but 3 is fine
    c <- 1
    c <- 4
		// c <- 8 : trying to add another element in the chanel will throw an error
    fmt.Println(<-c)
    fmt.Println(<-c)
}