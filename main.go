package main

import (
	"fmt"
)

func main() {
	var list LinkedList
	list = NewLinkedList(1)
	list.Append(2).
		Append(3).
		Append(4).
		Append(5).
		Append(6)

	fmt.Println(list)
	fmt.Println("---")

	list.Reverse()
	fmt.Println(list)
	fmt.Println("---")

	list.Reverse()
	fmt.Println(list)
}
