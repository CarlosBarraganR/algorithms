package main

import (
	"fmt"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge[T string | float32 | int](array []T, lo int, mid int, hi int) {
	aux := make([]T, len(array))

	for k := lo; k <= hi; k++ {
		aux[k] = array[k]
	}

	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i > mid {
			array[k] = aux[j]
			j++
		} else if j > hi {
			array[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			array[k] = aux[j]
			j++
		} else {
			array[k] = aux[i]
			i++
		}
	}
}

func MergesortBottomUp[T string | float32 | int](array []T) {
	n := len(array)
	for sz := 1; sz < n; sz = sz + sz {
		for lo := 0; lo < n-sz; lo += sz + sz {
			merge(array, lo, lo+sz-1, min(lo+sz+sz-1, n-1))
		}
	}
}

func main() {
	start := time.Now()

	arrayString := []string{"M", "E", "R", "G", "E", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	MergesortBottomUp(arrayString)
	fmt.Println(arrayString) // [A E E E E G L M M O P R R S T X]

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
