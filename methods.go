package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyInt int

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f *MyInt) Sum(num MyInt) {
	*f = MyInt(int(*f) + int(num))
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	pv := Vertex{3, 4}
	ScaleFunc(&pv, 10)

	var myNum MyInt = 10
	myNum.Sum(15)
	fmt.Println(myNum)

}
