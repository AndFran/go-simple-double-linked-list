package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value    int
	next     *Node
	previous *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.previous
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(value int) {
	n := &Node{value: value}
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.previous = l.tail
	}
	l.tail = n
}

func (l *List) Find(value int) (*Node, error) {
	if l.head == nil {
		return nil, errors.New("no data")
	}
	current := l.head
	for current != nil {
		if current.value == value {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("value not found")
}

func (l *List) Display() {
	fmt.Println("------------------")
	current := l.head
	for current != nil {
		fmt.Println(current)
		current = current.next
	}
	fmt.Println("------------------")
}

func (l *List) DisplayReverse() {
	fmt.Println("------------------")
	current := l.Last()
	for current != nil {
		fmt.Println(current)
		current = current.previous
	}

	fmt.Println("------------------")
}

func main() {
	ll := &List{}
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)
	ll.Push(4)
	ll.Push(5)
	ll.Display()
	ll.DisplayReverse()

	node, err := ll.Find(41)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data found in node: ", node)
	}
}
