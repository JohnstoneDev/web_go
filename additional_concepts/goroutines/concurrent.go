/*
	Program to illustrate multiple goroutines
*/


/*
	Goroutines can be used to run background operations in a program
	With goroutines, we can split one task in different segment to
	perform better
*/
package main

import (
	"fmt"
	"time"
)

func display(message string) {
	fmt.Println(message)
	fmt.Println("<<<<<<<<<<")
}

func main() {
	// run different goroutine
	go display("Process 1")
	go display("Process 2")
	go display("Process 3")

	// to sleep the main goroutine for 1 sec
	time.Sleep(time.Second * 1)
}