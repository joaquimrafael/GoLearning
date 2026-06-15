package main

import (
	"fmt"
	"math"
	"strconv"
)

type Individual interface {
	Hello()
	Grow()
}

type Pessoa struct {
	name string
	age  int
}

func (p *Pessoa) Hello() {
	if p == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Printf("Oi, meu nome e %s, tenho %d anos.\n", p.name, p.age)
}

func (p *Pessoa) Grow() {
	p.age++
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	p1 := Pessoa{"joaquim", 21}
	p2 := Pessoa{name: "rafael", age: 22}
	var p3 Pessoa

	p1.Hello()
	p2.Hello()
	p3.Hello()

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	p1.Grow()
	p1.Hello()

	var i Individual
	describe(i)

	i = &p2
	i.Hello()
	describe(i)

	var t *Pessoa
	i = t
	describe(i)
	i.Hello()

	fmt.Println()
	do(21)
	do("hello")
	do(true)
	fmt.Println()

	inteiro, err := strconv.Atoi("12")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", inteiro)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func describe(i Individual) {
	fmt.Printf("(%v, %T)\n", i, i)
}
