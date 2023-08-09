package main

/*
	Goroutines create concurrent programs, which are able to run
	multiple processes at a time
*/

import (
	"fmt"
)


// create a function
func display(message string) {
	fmt.Println(message)
}


func main() {
	// call / start goroutine (append go)
	go display("Process 1")

	display("Process 2")
}
