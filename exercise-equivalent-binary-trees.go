package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	// Test Walk
	{
		ch := make(chan int, 10)

		go Walk(tree.New(2), ch)
		fmt.Print("Walk(tree.New(1), ch) : ")
		for v := range ch {
			fmt.Print(v, " ")
		}
	}
}
