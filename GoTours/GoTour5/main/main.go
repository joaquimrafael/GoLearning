package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Add(val T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: val}
}

func (l *List[T]) Println() {
	current := l
	for current != nil {
		fmt.Println(*current)
		current = current.next
	}
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	var list List[int]
	list.Add(12)
	list.Add(11)
	list.Add(11)
	list.Add(11)
	list.Add(11)
	list.Println()

	var words List[string]
	words.Add("Hello")
	words.Add("World")
	words.Add("!")
	words.Println()

	var bytes = List[byte]{val: 12}
	bytes.Add(127)
	bytes.Println()
}
