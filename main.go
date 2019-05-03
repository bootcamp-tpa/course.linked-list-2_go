package main

import (
	"fmt"
	"strings"
)

// Node of linked list
type Node struct {
	Next  *Node
	Value interface{}
}

func (node *Node) String() string {
	return fmt.Sprintf("{ value: %v, next: %p } -> ", node.Value, node.Next)
}

// Append adds a node with the given value as the next node after this node.
func (node *Node) Append(value interface{}) *Node {
	nextNode := &Node{Value: value}
	node.Next = nextNode
	return nextNode
}

// LinkedList provides methods for manipulating a linked list.
type LinkedList struct {
	Head *Node
}

// NewLinkedList creates a linked list.
func NewLinkedList(headValue interface{}) LinkedList {
	head := Node{Value: headValue}
	return LinkedList{Head: &head}
}

// Append adds a node to the end of this list.
func (list *LinkedList) Append(value interface{}) *Node {
	node := &Node{Value: value}
	tail := list.FindTail()
	if tail == nil {
		list.Head = node
		return list.Head
	}
	tail.Next = node
	return tail.Next
}

// FindTail returns the tail node of this list.
func (list *LinkedList) FindTail() *Node {
	node := list.Head
	if node == nil {
		return nil
	}
	for node.Next != nil {
		node = node.Next
	}
	return node
}

// Reverse points all the "next" links of nodes in this list in the opposite direction.
func (list *LinkedList) Reverse() *Node {
	head := list.Head

	var prevNode *Node
	for head != nil {
		nextNode := head.Next
		head.Next = prevNode
		prevNode = head
		head = nextNode
	}

	list.Head = prevNode
	return list.Head
}

func (list LinkedList) String() string {
	var sb strings.Builder
	node := list.Head
	for node != nil {
		sb.WriteString(node.String())
		node = node.Next
	}
	return sb.String()
}

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
