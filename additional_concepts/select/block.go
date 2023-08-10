/*
	The select statement blocks all channels if they are not ready for execution
	A default case can be executed when none of the channels are ready
*/

package main

import (
	"fmt"
	"time"
)


func main() {
	// create two channels
	number := make(chan int)
	message := make(chan string)

	// function calls with goroutine
	go channelNumber(number)
	go channelMessage(message)

	// selects and executes a message
	select {
		case firstChannel := <- number :
			fmt.Println("Channel Data:", firstChannel)

		case secondChannel := <- message :
			fmt.Println("Channel Data:", secondChannel)

		// default case
		default:
			fmt.Println("Wait!! Channels are not ready for execution")
	}
}

// if we sleep both channels for the same time, the select will block both channels
// from execution for the duration slept then it will exeute one of the channels
// randomly because they will both be available

// goroutines that send data to channel
func channelNumber(number chan int) {
	// sleep the channel by a second
	time.Sleep(time.Second * 1)

	number <- 43
}


func channelMessage(message chan string) {
	// sleep the channel by a second
	time.Sleep(time.Second * 1)

	message <- "Hello there!"
}