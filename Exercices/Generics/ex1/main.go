package main

import "fmt"

type ValidTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func DoubleNumber[T ValidTypes](t T) T {
	return t * 2
}

func main() {
	fmt.Println(DoubleNumber(12))
	fmt.Println(DoubleNumber(51.12))
}
