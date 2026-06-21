package main

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (list *LinkedList[T]) Add(t T) {
	n := &Node[T]{
		Value: t,
	}

	if list.Head == nil {
		list.Head = n
		list.Tail = n
		return
	}

	list.Tail.Next = n
	list.Tail = list.Tail.Next
}

func (list *LinkedList[T]) Insert(t T, pos int) {
	n := &Node[T]{
		Value: t,
	}

	if list.Head == nil {
		list.Head = n
		list.Tail = n
		return
	}

	if pos <= 0 {
		n.Next = list.Head
		list.Head = n
		return
	}

	curNode := list.Head
	for i := 1; i < pos; i++ {
		if curNode.Next == nil {
			curNode.Next = n
			list.Tail = curNode.Next
			return
		}
		curNode = curNode.Next
	}
	n.Next = curNode.Next
	curNode.Next = n
	if list.Tail == curNode {
		list.Tail = n
	}
}

func (list *LinkedList[T]) Index(t T) int {
	i := 0
	for curNode := list.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == t {
			return i
		}
		i++
	}
	return -1
}

func main() {
	l := &LinkedList[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(300, 10)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Add(400)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(500, 6)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}
}
