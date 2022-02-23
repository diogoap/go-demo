package main

import (
	"fmt"

	"rsc.io/quote"
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	fmt.Println(quote.Go())

	//var name = "fff"
	//var name2 string = "test"

	age, name := 38, "Diogo"

	a, b := swap("Hello world", fmt.Sprintf("%s, %d years", name, age))
	fmt.Println(a, b)
	fmt.Println(sum(1, 5))

	v := false // change me!
	fmt.Printf("Big is of type %T\n", v)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func sum(a, b int) int {
	return a + b
}

func swap(x, y string) (first, last string) {
	first = y
	last = x
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
