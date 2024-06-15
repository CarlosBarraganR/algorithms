package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Shuffle[T int | float32 | string](array []T) {
	for i := 1; i < len(array); i++ {
		randomIndex := rand.Intn(i)
		array[i], array[randomIndex] = array[randomIndex], array[i]
	}
}

func main() {
	start := time.Now()
	arrayInt := []int{14, 6, 232, 5, 63, 2, 1, 5}
	fmt.Println(arrayInt) // [14 6 232 5 63 2 1 5]
	Shuffle(arrayInt)
	fmt.Println(arrayInt) // [6 232 1 5 14 5 2 63]

	arrayString := []string{"z", "b", "f", "a", "y", "e", "r", "p"}
	fmt.Println(arrayString) // [z b f a y e r p]
	Shuffle(arrayString)
	fmt.Println(arrayString) // [e f r y b a p z]

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
