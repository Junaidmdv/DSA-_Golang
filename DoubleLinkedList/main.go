package main

import (
	"fmt"
)

type Node struct {
	Data string
	Next *Node
	Prev *Node
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
}

func (list *DoubleLinkedList) PushBack(Data string) {
	NewNode := &Node{Data: Data, Next: nil}

	if list.Head == nil {
		list.Head = NewNode
		list.Tail = NewNode
	} else {
		list.Tail.Next = NewNode
		NewNode.Prev = list.Tail
		list.Tail = NewNode
	}
}

func (list *DoubleLinkedList) PushFront(Data string) {
	NewNode := &Node{Data: Data, Next: nil}
	if list.Head == nil {
		list.Head = NewNode
		list.Tail = NewNode
	} else {
		NewNode.Next = list.Head
		list.Head.Prev = NewNode
		list.Head = NewNode
	}
}

func (list *DoubleLinkedList) FindMidle() {
	fast := list.Head
	slow := list.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fmt.Println(slow.Data)

}

func (list *DoubleLinkedList) PopByIndex(Index int) {

}

func (list *DoubleLinkedList) PopBack() {

}

func (list *DoubleLinkedList) PopFront() {

}

func (list *DoubleLinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Printf("%s ", current.Data)
		current = current.Next
	}

	fmt.Println("nil")
}

func main() {
	linkedList := &DoubleLinkedList{}

	for i := range 26 {
		linkedList.PushFront(string(i + 97))
	}
	linkedList.Print()
	linkedList.FindMidle()

}
