package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var a, b int = 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	// pos, neg := adder(), adder()
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(
	// 		pos(i),
	// 		neg(-1*i),
	// 	)
	// }

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
