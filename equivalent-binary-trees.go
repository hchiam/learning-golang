// https://godoc.org/golang.org/x/tour/tree#Tree
// the tree package defines this Tree type:

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel c.
func Walk(t *tree.Tree, c chan int) {
	_walk(t, c)
	close(c)
}

func _walk(t *tree.Tree, c chan int) {
	if t != nil {
		_walk(t.Left, c)
		c <- t.Value
		_walk(t.Right, c)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for v := range c1 {
		if v != <-c2 {
			return false
		}
	}
	return true
}

func main() {
	c := make(chan int)
	go Walk(tree.New(1), c)
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("should be true:")
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println("should be false:")
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
