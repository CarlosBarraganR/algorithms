package main

import (
	"fmt"
	"time"
)

type BinaryHeap[T int | float32 | string] struct {
	items []T
}

func (bh *BinaryHeap[T]) sink(index int, n int) {
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

func (bh *BinaryHeap[T]) Sort() {
	heapSize := len(bh.items)
	// 1st sweep for internal tree nodes sort
	for index := heapSize/2 - 1; index >= 0; index-- {
		bh.sink(index, heapSize)
	}
	// 2nd sweep for swapping root [0] for [heapSize - 1] until heap is reduced to root index
	for heapSize > 1 {
		bh.items[0], bh.items[heapSize-1] = bh.items[heapSize-1], bh.items[0]
		// reduce heap size each time so sinks stops on [heapSize - 1]
		heapSize--           // otherwise sink will overflow to indeces outside the heapSize
		bh.sink(0, heapSize) // this sink is different because it's needed to stop
	}

}

func main() {
	start := time.Now()

	array := &BinaryHeap[string]{[]string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}}

	fmt.Println(array.items)
	array.Sort()
	fmt.Println(array.items)

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
