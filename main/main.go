package main

import (
	"doubleLinkedList"
	"fmt"
)

func main() {
	List := doubleLinkedList.Node{
		Value: 0,
	}

	fmt.Println(List.GetValue(12))

	List.AddNode(&doubleLinkedList.Node{Value: 1})
	List.AddNode(&doubleLinkedList.Node{Value: 2})
	List.AddNode(&doubleLinkedList.Node{Value: 3})
	List.AddNode(&doubleLinkedList.Node{Value: 4})

	fmt.Println("Tous les éléments de ma liste")
	List.TraverseList()

	fmt.Println("L'index 2")
	fmt.Println(List.GetValue(2))

	err := List.DeleteNode(2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("L'index 2 après avoir supprimé l'ancien index 2")
	fmt.Println(List.GetValue(2))

	fmt.Println("Tous les éléments de ma nouvelle liste")
	List.TraverseList()
}
