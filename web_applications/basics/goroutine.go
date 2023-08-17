package main

/*
	NOTES :
	Don't use shared data to communicate; use communication to share data
	In Go 1.5,the runtime now sets the default number of threads to run simultaneously, defined by GOMAXPROCS,
 	to the number of cores available on the CPU.*/

import (
	"fmt"
	"runtime"
)

// define a function
func say(s string) {
	for i := 0; i  < 5; i++ {
		runtime.Gosched() // let the CPU execute other goroutines & come back at some point
		fmt.Println(s)
	}
}

func main() {
	go say("the way") // creata a new routine
	say("This is") // current goroutine
}