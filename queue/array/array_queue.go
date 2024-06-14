package main

import (
	"fmt"
	"time"
)

type Queue[T any] struct {
	items []T
}

func (queue *Queue[T]) enqueue(item T) {
	queue.items = append(queue.items, item)
}

func (queue *Queue[T]) dequeue() T {
	dequeuedValue := queue.items[0]
	queue.items = queue.items[1:len(queue.items)]
	return dequeuedValue
}

func (queue *Queue[T]) print() {
	fmt.Printf("queue.items: %v\n", queue.items)
}

func main() {
	start := time.Now()

	queue := &Queue[int]{}

	queue.print() // []
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.print()                //  [1 2 3]
	fmt.Println(queue.dequeue()) // 1
	fmt.Println(queue.dequeue()) // 2
	queue.print()                //  [3]

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
