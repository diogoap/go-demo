package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, name string) {
	if t != nil {
		Walk(t.Left, ch, name)
		ch <- t.Value
		println(name, ": ", t.Value)
		Walk(t.Right, ch, name)
	}
}

func goWalk(t *tree.Tree, ch chan int, name string) {
	Walk(t, ch, name)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int, 10), make(chan int, 10)

	go goWalk(t1, c1, "walk1")
	go goWalk(t2, c2, "walk2")

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	//fmt.Println(Same(tree.New(1), tree.New(2)))
	//fmt.Println(Same(tree.New(2), tree.New(2)))
}
