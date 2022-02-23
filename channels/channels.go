package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main1() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 3)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	go sum(s, c)

	for val := range c {
		fmt.Println(val)
	}

	x, y, z := <-c, <-c, <-c // receive from c

	fmt.Println(x, y, z, x+y+z)

	///buffered channel
	ch := make(chan int, 30)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//Range and close
	fmt.Println("Fibonacci")

	chan2 := make(chan int, 15)
	go fibonacci1(cap(chan2), chan2)
	for i := range chan2 {
		fmt.Println(i)
	}
}
