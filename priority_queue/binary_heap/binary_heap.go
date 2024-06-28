package main

import (
	"fmt"
	"time"
)

type BinaryHeap[T int | float32 | string] struct {
	items []T
}

func (pq *BinaryHeap[T]) swim(index int) {
	// calculate parent from tree
	parent := (index / 2) - 1
	// swim up until the top or minor than parent
	for parent > 0 {
		// if parent is minor than current index, swap index for parent
		if pq.items[parent] < pq.items[index] {
			pq.items[parent], pq.items[index] = pq.items[index], pq.items[parent]
		}
		index = parent            // move up to parent position
		parent = (parent / 2) - 1 // find next parent index
	}
}

func (pq *BinaryHeap[T]) sink(index int) {
	n := len(pq.items)
	// Swim down while the element has at least one child
	for (2*index + 1) < n {
		j := 2*index + 1 // left child index

		// if right child exists and is greater than left, move to right child index
		if j+1 < n && pq.items[j] < pq.items[j+1] {
			j++
		}
		// if the current element is greater or equal than largest child, stop the sinking because index is in place
		if pq.items[index] >= pq.items[j] {
			break
		}
		// Swap the current element with the largest child
		pq.items[index], pq.items[j] = pq.items[j], pq.items[index]
		index = j // move down to the swaped child position
	}
}

func (pq *BinaryHeap[T]) IsEmpty() bool {
	return len(pq.items) == 0
}

func (pq *BinaryHeap[T]) Push(value T) {
	pq.items = append(pq.items, value)
	pq.swim(len(pq.items) - 1)
}

func (pq *BinaryHeap[T]) Pop() T {
	// store next in queue
	valueToPop := pq.items[0]
	// exchange root [0] for last [n-1]
	pq.items[0], pq.items[len(pq.items)-1] = pq.items[len(pq.items)-1], pq.items[0]
	// slice to remove last item
	pq.items = pq.items[:len(pq.items)-1]
	// sink switched root until correct position
	pq.sink(0)

	return valueToPop
}

func main() {
	start := time.Now()

	priorityQueue := &BinaryHeap[string]{[]string{"T", "P", "R", "N", "H", "O", "A", "E", "I", "G"}}

	fmt.Println(priorityQueue.IsEmpty()) // false
	fmt.Println(priorityQueue.items)     // [T P R N H O A E I G S]

	priorityQueue.Push("S")
	fmt.Println(priorityQueue.items) // [T S R N P O A E I G H]

	fmt.Println(priorityQueue.Pop()) // T
	fmt.Println(priorityQueue.items) // [S P R N H O A E I G]

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
