/*
	Channels in go act as a medium for goroutines to communicate with each other
	We use channles to allow goroutines to communicate & share resources
*/

package main

import "fmt"

func main() {
	// create a channel of integer type
	number := make(chan int)

	// access type and value of channel
	fmt.Printf("Channel Type: %T\n", number)
	fmt.Printf("Channel Type: %v\n", number) // value of a channel is a memory address

}