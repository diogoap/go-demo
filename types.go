package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

type MyInt2 int

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{Y: 1}  // X:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)

	//var my MyInt2 = 2 //OK
	my := MyInt2(2)
	fmt.Println(my)
}
