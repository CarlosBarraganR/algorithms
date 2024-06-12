package main

import (
	"fmt"
	"time"
)

type Node struct {
	value string
	next  *Node
}

type LinkedList struct {
	node *Node
}

func (list *LinkedList) print() {
	current := list.node
	// we need to iterate to the end of the LL to push
	for current != nil {
		fmt.Printf("%s => ", current.value)
		current = current.next
	}
	fmt.Println(nil)
}

func (list *LinkedList) push(value string) {
	node := &Node{value: value}
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

func (list *LinkedList) pop() string {
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
	linkedList := &LinkedList{}
	linkedList.push("hello")
	linkedList.push("my")
	linkedList.print()            // hello => my => <nil>
	fmt.Println(linkedList.pop()) // my
	linkedList.print()            // hello => <nil>

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
