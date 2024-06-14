package main

import (
	"fmt"
	"time"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	first *Node[T]
	last  *Node[T]
}

func (list *LinkedList[T]) isEmpty() bool {
	return list.first == nil
}

func (list *LinkedList[T]) enqueue(value T) {
	newNode := &Node[T]{value: value}
	if list.isEmpty() {
		list.first = newNode
		list.last = newNode
	} else {
		list.last.next = newNode
		list.last = newNode
	}
}

func (list *LinkedList[T]) dequeue() T {
	// TODO: guard for when the LinkedList is empty to avoid null poiner

	dequeuedValue := list.first.value
	list.first = list.first.next

	return dequeuedValue
}

func (list *LinkedList[T]) print() {
	current := list.first
	for current != nil {
		fmt.Printf("%v => ", current.value)
		current = current.next
	}
	fmt.Println(nil)
}

func main() {
	start := time.Now()

	queueInt := &LinkedList[int]{}
	fmt.Println(queueInt.isEmpty()) // true
	queueInt.enqueue(1)
	fmt.Println(queueInt.isEmpty()) // false
	queueInt.enqueue(2)
	queueInt.enqueue(3)
	queueInt.enqueue(5)
	queueInt.print()                // 1 => 2 => 3 => 5 => <nil>
	fmt.Println(queueInt.dequeue()) // 1
	fmt.Println(queueInt.dequeue()) // 2
	queueInt.print()                // 3 => 5 => <nil>

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)

}
