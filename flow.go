package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var previous float64

	for i := 1.0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)

		if previous == z {
			fmt.Println("Stopping at attempt: ", i)
			return z
		}

		previous = z
	}

	return z
}

func main() {
	fmt.Println(Sqrt(3))

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
