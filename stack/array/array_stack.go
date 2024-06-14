package main

import (
	"fmt"
	"time"
)

type Stack[T any] struct {
	items []T
}

func (stack *Stack[T]) push(item T) {
	stack.items = append(stack.items, item)
}

func (stack *Stack[T]) pop() T {
	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return lastItem
}

func (stack *Stack[T]) print() {
	fmt.Printf("%v\n", stack.items)
}

func (stack *Stack[T]) isEmpty() bool {
	return len(stack.items) == 0
}

func main() {
	start := time.Now()

	stackStrings := &Stack[string]{}
	fmt.Println(stackStrings.isEmpty()) // true
	stackStrings.push("Hello")
	stackStrings.push("World")
	stackStrings.push("My")
	stackStrings.push("Test")
	stackStrings.print()                // [Hello World My Test]
	fmt.Println(stackStrings.pop())     // Test
	stackStrings.print()                // [Hello World My]
	fmt.Println(stackStrings.isEmpty()) // false

	stackInts := &Stack[int]{}
	stackInts.push(2)
	stackInts.push(10)
	stackInts.push(200)
	stackInts.push(23)
	stackInts.print()            // [2 10 200 23]
	fmt.Println(stackInts.pop()) // 23
	stackInts.print()            // [2 10 200]

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)

}
