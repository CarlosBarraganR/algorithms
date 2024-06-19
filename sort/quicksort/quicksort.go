package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle[T int | float32 | string](array []T) {
	for i := 1; i < len(array); i++ {
		randomIndex := rand.Intn(i)
		array[i], array[randomIndex] = array[randomIndex], array[i]
	}
}

func partition[T int | float32 | string](array []T) int {
	pivot := array[0]
	i := 1

	for j := 1; j < len(array); j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}

	// final move to set pivot to the index it belongs
	array[0], array[i-1] = array[i-1], array[0]
	// return pivot correct index
	return i - 1
}

func Quicksort[T int | float32 | string](array []T) {
	// recursive base condition
	if len(array) <= 1 {
		return
	}

	shuffle(array) // shuffle the array to performance gains

	pivotIndex := partition(array)
	Quicksort(array[:pivotIndex])   // sort left side
	Quicksort(array[pivotIndex+1:]) // sort right side

}

func main() {
	start := time.Now()

	arrayInt := []int{14, 6, 232, 5, 63, 2, 1, 5}
	Quicksort(arrayInt)
	fmt.Println(arrayInt) // [1 2 5 5 6 14 63 232]

	arrayString := []string{"Q", "U", "I", "C", "K", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	Quicksort(arrayString)
	fmt.Println(arrayString) // [A C E E I K L M O P Q R S T U X]

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
