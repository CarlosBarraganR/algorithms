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
	node *Node[T]
}

func (list *LinkedList[T]) print() {
	current := list.node
	// we need to iterate to the end of the LL to push
	for current != nil {
		fmt.Printf("%v => ", current.value)
		current = current.next
	}
	fmt.Println(nil)
}

func (list *LinkedList[T]) push(value T) {
	node := &Node[T]{value: value}
	// if node is null it means it's the first push of the LL
	if list.node == nil {
		list.node = node
	} else {
		current := list.node
		// we need to iterate to the end of the LL to push
		for current.next != nil {
			current = current.next
		}
		// then add the new value to the end
		current.next = node
	}
}

func (list *LinkedList[T]) pop() T {
	current := list.node
	// we iterate to the one before the last item in LL
	for current.next.next != nil {
		current = current.next
	}
	// buffer var for popedValue
	popedValue := current.next.value
	// poped index goes null
	current.next = nil

	return popedValue

}

func main() {
	start := time.Now()
	linkedListStrings := &LinkedList[string]{}
	linkedListStrings.push("hello")
	linkedListStrings.push("my")
	linkedListStrings.print()            // hello => my => <nil>
	fmt.Println(linkedListStrings.pop()) // my
	linkedListStrings.print()            // hello => <nil>

	linkedListInts := &LinkedList[int]{}

	linkedListInts.push(1)
	linkedListInts.push(2)
	linkedListInts.push(4)
	linkedListInts.print()
	fmt.Println(linkedListInts.pop())
	linkedListInts.print()

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
