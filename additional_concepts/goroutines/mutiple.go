/*
	Program to run two goroutines concurrently
*/

package main

import (
	"fmt"
	"time"
)

// create a function (to be used as routine)
func display(message string) {
	// infinite loop
	for {
		fmt.Println(message)

		// sleep for one second'
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// call the function with go routine
	go display("Process 1")
	display("Process 2")
}