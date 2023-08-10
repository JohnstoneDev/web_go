/*
	A select statement allows for execution of a channel among many alternatives
	The syntax of select case looks similar to that of the Switch Case in Go.
	And, like the switch case, only one of the cases is executed by select.
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
	}
}


// goroutines that send data to channel
func channelNumber(number chan int) {
	number <- 23
}

// if we sleep one channel, the other is always ready for
// execution & will be auto picked by the select

func channelMessage(message chan string) {
	// sleep the channel by a second
	time.Sleep(time.Second * 1)

	message <- "Hello there!"
}