package main

import (
	"fmt"
	"sync"
)

type Person struct {
	FirstName, LastName string
	Age                 int
}

func MakePerson(firstName, lastName string, age int) Person {
	person := Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
	return person
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	person := Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
	return &person
}

func main() {
	p1 := MakePerson("Olavo", "Neto", 20)
	p2 := MakePersonPointer("Joaquim", "Prieto", 21)
	fmt.Println(p1, p2)

	var pointer *int
	var function func(a, b int) bool

	function = func(n1, n2 int) bool {
		if n1 > n2 {
			return true
		}
		return false
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("hello goroutine")

	wg.Wait()

	fmt.Println(pointer)
	fmt.Println(function(12, 3))
}
