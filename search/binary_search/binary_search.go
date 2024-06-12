package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	array := make([]int, 10)
	for i := 0; i < len(array); i++ {
		array[i] = i
	}
	fmt.Println(BinarySearch(array, 2))  // 2
	fmt.Println(BinarySearch(array, 20)) // -1

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}

/*
 * Binary Search
 * Complexity: Olog
 */
func BinarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := left + (right-left)/2

		// found it
		if array[mid] == target {
			return mid
		}

		// move mid to keep iteration
		if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // not found
}
