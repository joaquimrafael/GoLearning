package main

import (
	"fmt"
	"log"

	"github.com/joaquimrafael/golearning/testsgo/calculator"
	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println(calculator.Add(decimal.NewFromInt(5), decimal.NewFromInt(10)))
	log.Fatal("hi")
}
