package main

import (
	"fmt"
	"math"
	"runtime"
)

// defer: "world" so imprime depois do "hello" (quando hello retorna)
func hello() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// defers empilham (LIFO): imprime 9, 8, ... 0 ao final
func counting() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// switch com atribuicao no inicio (os := runtime.GOOS)
func printOS() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

// Sqrt: raiz quadrada pelo metodo de Newton
func Sqrt(x float64) float64 {
	z := 1.0
	var y float64
	for math.Abs(z-y) > 1e-10 {
		y = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	// --- for classico (init; condition; post) ---
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	// --- for como "while" (so a condicao) ---
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// --- if/else ---
	fmt.Println("2 e par? ", isEven(2))
	fmt.Println("3 e par? ", isEven(3))

	// --- comparando minha Sqrt com a da stdlib ---
	fmt.Println()
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println()

	// --- switch e defers ---
	printOS()
	hello()
	counting()
}
