package main

import (
	"fmt"
	"time"
)

/*
 * Weighted Quick Union
 * Complexity: Olog
 */
func main() {
	start := time.Now()
	array := make([]int, 10)
	size := make([]int, 10)

	for i := 0; i < 10; i++ {
		array[i] = i
		size[i] = 1 // add initial weigth size to all nodes
	}

	fmt.Println(array) // [0 1 2 3 4 5 6 7 8 9]

	union(4, 3, &array, &size)
	union(3, 8, &array, &size)
	union(6, 5, &array, &size)
	union(9, 4, &array, &size)
	union(2, 1, &array, &size)
	union(5, 0, &array, &size)
	union(7, 2, &array, &size)
	union(6, 1, &array, &size)
	union(6, 1, &array, &size)
	union(7, 3, &array, &size)

	fmt.Println(connected(4, 9, &array)) // true
	fmt.Println(array)                   // [6 2 6 4 6 6 6 2 4 4]

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}

func root(p int, array *[]int) int {
	if (*array)[p] != p {
		return root((*array)[p], array)
	}
	return (*array)[p]
}

func union(p int, q int, array *[]int, size *[]int) {
	var i = root(p, array)
	var j = root(q, array)

	// if same root they are already united
	if i == j {
		return
	}
	// we weight wich tree is smaller
	if (*size)[i] < (*size)[j] {
		(*array)[i] = j
		(*size)[j] += (*size)[i]
	} else {
		(*array)[j] = i
		(*size)[i] += (*size)[j]
	}
}

func connected(p int, q int, array *[]int) bool {
	return root(p, array) == root(q, array)
}
