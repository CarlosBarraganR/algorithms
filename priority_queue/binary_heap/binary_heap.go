package main

import (
	"fmt"
	"time"
)

type BinaryHeap[T int | float32 | string] struct {
	items []T
}

func (bh *BinaryHeap[T]) swim(index int) {
	// calculate parent from tree
	parent := (index / 2) - 1
	// swim up until the top or minor than parent
	for parent > 0 {
		// if parent is minor than current index, swap index for parent
		if bh.items[parent] < bh.items[index] {
			bh.items[parent], bh.items[index] = bh.items[index], bh.items[parent]
		}
		index = parent            // move up to parent position
		parent = (parent / 2) - 1 // find next parent index
	}
}

func (bh *BinaryHeap[T]) sink(index int) {
	n := len(bh.items)
	// Swim down while the element has at least one child
	for (2*index + 1) < n {
		j := 2*index + 1 // left child index

		// if right child exists and is greater than left, move to right child index
		if j+1 < n && bh.items[j] < bh.items[j+1] {
			j++
		}
		// if the current element is greater or equal than largest child, stop the sinking because index is in place
		if bh.items[index] >= bh.items[j] {
			break
		}
		// Swap the current element with the largest child
		bh.items[index], bh.items[j] = bh.items[j], bh.items[index]
		index = j // move down to the swaped child position
	}
}

func (bh *BinaryHeap[T]) IsEmpty() bool {
	return len(bh.items) == 0
}

func (bh *BinaryHeap[T]) Push(value T) {
	bh.items = append(bh.items, value)
	bh.swim(len(bh.items) - 1)
}

func (bh *BinaryHeap[T]) Pop() T {
	// store next in queue
	valueToPop := bh.items[0]
	// exchange root [0] for last [n-1]
	bh.items[0], bh.items[len(bh.items)-1] = bh.items[len(bh.items)-1], bh.items[0]
	// slice to remove last item
	bh.items = bh.items[:len(bh.items)-1]
	// sink switched root until correct position
	bh.sink(0)

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
