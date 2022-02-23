package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt2(x float64) (float64, error) {

	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := float64(1)
	var previous float64

	for i := 1.0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)

		if previous == z {
			return z, nil
		}

		previous = z
	}

	return z, nil
}

func main() {
	print(2)
	print(-2)
}

func print(num float64) {
	if res, err := Sqrt2(num); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println("ERROR: ", err)
	}
}
