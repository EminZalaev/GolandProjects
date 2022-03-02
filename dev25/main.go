//Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"
)

func mySleep(ch <-chan string, d time.Duration) {

	// Select statement
	select {

	// Using case statement to receive
	// or send operation on channel
	case output := <-ch:
		fmt.Println(output)

	// Calling After() method with its
	// parameter
	case <-time.After(d):

		// Printed after 5 seconds
		fmt.Println("Its timeout..")
	}
}

func main() {
	// Creating a channel
	// Using make keyword
	channel := make(chan string, 2)
	fmt.Println("Start wait")
	mySleep(channel, time.Second*3)
}
