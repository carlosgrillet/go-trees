package main

import "fmt"

func main() {
	numbers := NewTree()
	fmt.Printf("Tree max depth: %d\n", numbers.MaxDepth())
	fmt.Println(numbers)
}
