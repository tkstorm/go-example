package main

import (
	"fmt"
	"unicode"
)

func squares(c chan interface{}) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	for _, r := range "Hey, man!" {
		c <- unicode.ToUpper(r)
	}

	close(c) // close channel
}

func main() {
	fmt.Println("main() started")
	c := make(chan interface{})

	go squares(c) // start goroutine

	// periodic block/unblock of main goroutine until chanel closes
	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "<-- loop broke!")
			break // exit break loop
		} else {
			fmt.Println(val, ok)
		}
	}

	fmt.Println("main() stopped")
}
