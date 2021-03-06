package main

import "fmt"

type I interface {
	M()
}

type Empty interface {
	N()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	var generic interface{}
	describe(generic)

	var empty Empty
	describe(empty)
	empty.N()
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
