package main

import (
	"fmt"
	"time"
)

/*
 * Binary Search
 * Complexity: Olog
 */
func BinarySearch[T int | float32 | string](array []T, target T) int {
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

func main() {
	start := time.Now()

	arrayInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(BinarySearch(arrayInts, 2))  // 2
	fmt.Println(BinarySearch(arrayInts, 20)) // -1

	arrayStrings := []string{"apple", "banana", "cherry", "date", "fig", "grape"}
	fmt.Println(BinarySearch(arrayStrings, "cherry"))     // 2
	fmt.Println(BinarySearch(arrayStrings, "strawberry")) // -1

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
