package main

import (
	"fmt"
	"time"
)

func Sort[T int | float32 | string](array []T) {
	for i := range array {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func main() {
	start := time.Now()
	arrayInt := []int{14, 6, 232, 5, 63, 2, 1, 5}
	fmt.Println(arrayInt) // [14 6 232 5 63 2 1 5]
	Sort(arrayInt)
	fmt.Println(arrayInt) // [1 2 5 5 6 14 63 232]

	arrayString := []string{"z", "b", "f", "a", "y", "e", "r", "p"}
	fmt.Println(arrayString) // [z b f a y e r p]
	Sort(arrayString)
	fmt.Println(arrayString) // [a b e f p r y z]

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
