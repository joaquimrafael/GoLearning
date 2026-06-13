package main

import "fmt"

func fibonacci() func() int {
	n1, n2 := 0, 1
	return func() int {
		fib := n1
		n1, n2 = n2, n1+n2
		return fib
	}
}

func main() {
	f := fibonacci()
	s := make([]int, 10)
	for range s {
		fmt.Println(f())
	}
}
