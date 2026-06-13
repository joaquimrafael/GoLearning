package main

import (
	bib "GoTour1/bib"
	"fmt"
	"math"
	"math/cmplx"
)

var legal, bonito bool

var i, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const pi = 3.14

func main() {
	fmt.Printf("A soma de 76 e 24 e: %d\n", bib.Add(76, 24))
	resultado := bib.Add(-1, 21)
	fmt.Println(resultado)
	resultado = bib.Subtract(12, 8)
	fmt.Println(resultado)

	a, b := bib.Swap("hello", "world")
	println(a, b)

	fmt.Println(bib.Split(17))
	fmt.Println()

	println(i, bonito)

	var c, python, java = true, false, "no!"
	k := 3
	fmt.Println(i, j, c, python, java, k)

	fmt.Println()
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println()
	var initial int
	var f float64
	var bo bool
	var s string
	fmt.Printf("%v %v %v %q\n", initial, f, bo, s)

	fmt.Println()
	var x, y int = 3, 4
	var fl float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(fl)
	fmt.Println(x, y, z)

	const name = "Joaquim"
	fmt.Printf("\nOla %s o Pi sempre eh %.2f\n", name, pi)
}
