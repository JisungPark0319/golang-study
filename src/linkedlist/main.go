package main

import "dataStruct"

func main() {
	list := &dataStruct.LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()
	list.PrintReverse()

	list.RemoveNode(list.Root.Next)
	list.PrintNodes()
}
