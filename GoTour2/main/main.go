package main

import (
	"fmt"
	"math"
	"runtime"
)

func hello() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func counting() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func os() {
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
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println("2 e par? ", isEven(2))
	fmt.Println("3 e par? ", isEven(3))

	fmt.Println()
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println()

	os()
	hello()
	counting()
}
