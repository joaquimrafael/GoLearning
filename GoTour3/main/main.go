package main

import (
	"fmt"
)

func main() {
	i, j := 20, 11
	var k, n int
	var p *int
	p2 := &i
	p = &j

	fmt.Println(k, float64(n))
	fmt.Println(*p)
	fmt.Println(*p2)

}
