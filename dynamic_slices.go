package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 15)
	printSlice("b", b)

	var s = []int{6}
	printSlice("s", s)

	s = append(s, 10)
	printSlice("s", s)

	s = append(s, 15, 20, 25, 40)
	printSlice("s", s)

	for i, v := range s {
		fmt.Printf("position %d = %d\n", i, v)
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
