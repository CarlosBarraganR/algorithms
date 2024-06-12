package main

import (
	"fmt"
	"time"
)

type Stack struct {
	items []string
}

func (stack *Stack) push(item string) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) pop() string {
	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return lastItem
}

func (stack *Stack) print() {
	fmt.Printf("%s\n", stack.items)
}

func (stack *Stack) isEmpty() bool {
	return len(stack.items) == 0
}

func main() {
	start := time.Now()

	stack := &Stack{}
	fmt.Println(stack.isEmpty()) // true
	stack.push("Hello")
	stack.push("World")
	stack.push("My")
	stack.push("Test")
	stack.print()                // [Hello World My Test]
	fmt.Println(stack.pop())     // Test
	stack.print()                // [Hello World My]
	fmt.Println(stack.isEmpty()) // false

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)

}
