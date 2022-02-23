package main

import (
	"fmt"
	"time"
)

func main() {

	bomb()

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	//ch <- 3

	//tryReadRange(ch)

	//tryRead(ch)
	//tryRead(ch)
	//tryRead(ch)
	//tryRead(ch)
}

func tryRead(c chan int) {
	if val, isOpen := <-c; isOpen {
		fmt.Println("value: ", val, "is isOpen: ", isOpen)
	} else {
		fmt.Println("closed")
	}
}

func tryReadRange(c chan int) {
	items := cap(c)

	fmt.Println("capacity: ", items)

	for i := 0; i < items; i++ {
		fmt.Println(<-c)
	}
}

func bomb() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
