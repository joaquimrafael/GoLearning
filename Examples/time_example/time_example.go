package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("The time now is: ", time.Now())

	year, month, day := time.Now().Date()

	fmt.Printf("%d of %s of %d\n", day, month, year)
}
