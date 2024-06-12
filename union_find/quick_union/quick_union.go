package main

import (
	"fmt"
	"time"
)

/*
 * Quick Union (lazy)
 * Complexity: ON
 */
func main() {
	start := time.Now()
	array := make([]int, 10)

	for i := 0; i < 10; i++ {
		array[i] = i
	}

	fmt.Println(array) // [0 1 2 3 4 5 6 7 8 9]

	union(3, 8, &array)
	union(4, 3, &array)
	union(6, 5, &array)

	fmt.Println(connected(4, 9, &array)) // false
	fmt.Println(connected(3, 8, &array)) // true
	fmt.Println(connected(4, 8, &array)) // true

	fmt.Println(array) // [0 1 2 8 3 6 6 7 8 9]

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}

func root(p int, array *[]int) int {
	if (*array)[p] != p {
		return root((*array)[p], array)
	}
	return (*array)[p]
}

func union(p int, q int, array *[]int) {
	(*array)[root(p, array)] = q
}

func connected(p int, q int, array *[]int) bool {
	return root(p, array) == root(q, array)
}
