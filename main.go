package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	c := add(42, 42)
	fmt.Printf("c = a + b = %d\n", c)
	fmt.Println("Hello world,no ; needed !!!")
}
