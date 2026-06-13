package main

import (
	bib "GoTour1/bib"
	"fmt"
	"math"
	"math/cmplx"
)

// --- Variaveis de pacote ---

var legal, bonito bool // sem valor: ambas false

var i, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const pi = 3.14

func main() {
	// --- usando funcoes do pacote bib ---
	fmt.Printf("A soma de 76 e 24 e: %d\n", bib.Add(76, 24))
	resultado := bib.Add(-1, 21)
	fmt.Println(resultado)
	resultado = bib.Subtract(12, 8)
	fmt.Println(resultado)

	a, b := bib.Swap("hello", "world") // multiplos retornos
	println(a, b)

	fmt.Println(bib.Split(17))
	fmt.Println()

	println(i, bonito)

	// --- declaracao de variaveis ---
	var c, python, java = true, false, "no!" // tipos inferidos
	k := 3                                   // short declaration
	fmt.Println(i, j, c, python, java, k)

	// --- %T mostra o tipo, %v o valor ---
	fmt.Println()
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// --- valores zero (int, float, bool, string) ---
	fmt.Println()
	var initial int
	var f float64
	var bo bool
	var s string
	fmt.Printf("%v %v %v %q\n", initial, f, bo, s)

	// --- conversao de tipo (int -> float64 -> uint) ---
	fmt.Println()
	var x, y int = 3, 4
	var fl float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(fl)
	fmt.Println(x, y, z)

	// --- constante local ---
	const name = "Joaquim"
	fmt.Printf("\nOla %s o Pi sempre eh %.2f\n", name, pi)
}
