package main

import (
	"container/list"
	"fmt"
	"golang.org/x/tour/tree"
	"sort"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	fmt.Println("Walking", t)
	queue := list.New()
	queue.PushBack(t)
	for queue.Len() > 0 {
		current := queue.Back()
		queue.Remove(current)
		var currentTree = current.Value.(*tree.Tree)

		ch <- currentTree.Value

		if currentTree.Left != nil {
			queue.PushBack(currentTree.Left)
		}

		if currentTree.Right != nil {
			queue.PushBack(currentTree.Right)
		}
	}
	fmt.Println("Finished Walking", t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chan1 := make(chan int, 10)
	chan2 := make(chan int, 10)
	defer close(chan1)
	defer close(chan2)

	go Walk(t1, chan1)
	go Walk(t2, chan2)

	arr1 := make([]int, 10)
	arr2 := make([]int, 10)

	for i := 0; i < 10; i++ {
		arr1[i] = <-chan1
	}

	for i := 0; i < 10; i++ {
		arr2[i] = <-chan2
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i := 0; i < 10; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func main() {
	tree10 := tree.New(1)
	anotherTree10 := tree.New(1)
	tree20 := tree.New(2)

	fmt.Println(tree10, " == ", anotherTree10, " true ==", Same(tree10, anotherTree10))
	fmt.Println(tree10, " == ", tree20, "false ==", Same(tree10, tree20))
}
