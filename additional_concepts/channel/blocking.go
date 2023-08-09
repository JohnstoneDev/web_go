/*
	Channel automatically blocks the send and receive operations
	depending on the goroutines

	When a goroutine receives data from a channel, the operation is blocked
	unitl another goroutine sends the data to the channel
*/

package main

import "fmt"

func sendData(ch chan string) {
	// data sent to the channel
	ch <- "Received. Send Operation Successful"

	fmt.Println("No receiver! Send Operation Blocked!")
}

func main() {
	// create a channel
	ch := make(chan string)

	// function call with goroutine
	go sendData(ch)

	// receive channel data
	receiveData(ch)
}

func receiveData(ch chan string) {
	// receive data from the chanel
	fmt.Println(<- ch)
}
