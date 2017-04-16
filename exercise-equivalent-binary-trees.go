package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func TreeWalker(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		TreeWalker(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		TreeWalker(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	TreeWalker(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	// Test Walk
	{
		ch := make(chan int, 10)

		go Walk(tree.New(1), ch)
		fmt.Print("Walk(tree.New(1), ch) : ")
		for v := range ch {
			fmt.Print(v, " ")
		}
		fmt.Print("\n")
	}

	// Test Same
	{
		fmt.Println("Same(tree.New(1), tree.New(1)) : ",
			Same(tree.New(1), tree.New(1)))
		fmt.Println("Same(tree.New(1), tree.New(2)) : ",
			Same(tree.New(1), tree.New(2)))
	}
}
