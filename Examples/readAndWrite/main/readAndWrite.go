package main

import "fmt"

func main() {
	var (
		name string
		age  int
	)

	fmt.Println("Type your first name and age (separeted by space):")
	_, err := fmt.Scan(&name, &age)
	if err != nil {
		fmt.Println("Error reading input", err)
		return
	}
	fmt.Printf("Hello %s!\n", name)
	if age >= 18 {
		fmt.Println("You're an adult!")
	} else if age > 12 {
		fmt.Println("You're just a teenager.")
	} else {
		fmt.Println("Hey kid!")
	}
}
