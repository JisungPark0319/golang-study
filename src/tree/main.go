package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	tree := dataStruct.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}

	tree.DFS1()
	fmt.Println()

	tree.DFS2()
	fmt.Println()

	tree.BFS()
	fmt.Println()

	fmt.Println("********Binary Tree********")
	// Binary Tree
	btree := dataStruct.NewBinaryTree(5)
	btree.Root.AddNode(3)
	btree.Root.AddNode(2)
	btree.Root.AddNode(4)
	btree.Root.AddNode(8)
	btree.Root.AddNode(7)
	btree.Root.AddNode(6)
	btree.Root.AddNode(10)
	btree.Root.AddNode(9)

	btree.Print()

	fmt.Println()
	foundVal := 6
	if found, cnt := btree.Search(foundVal); found {
		fmt.Println("Found ", foundVal, " cnt: ", cnt)
	} else {
		fmt.Println("Not Found ", foundVal, " cnt: ", cnt)
	}
}
