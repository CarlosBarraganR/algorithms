package main

import (
	"fmt"
	"time"
)

func merge[T string | float32 | int](array []T, aux []T, lo int, mid int, hi int) {
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

func sort[T string | float32 | int](array []T, aux []T, lo int, hi int) {
	mid := lo + (hi-lo)/2

	// recursive base condition
	if hi <= lo {
		return
	}

	sort(array, aux, lo, mid)
	sort(array, aux, mid+1, hi)
	merge(array, aux, lo, mid, hi)
}

func Mergesort[T string | float32 | int](array []T) {
	aux := make([]T, len(array))
	sort(array, aux, 0, len(array)-1)
}

func main() {
	start := time.Now()

	arrayInt := []int{14, 6, 232, 5, 63, 2, 1, 5}
	Mergesort(arrayInt)
	fmt.Println(arrayInt)

	arrayString := []string{"z", "b", "f", "a", "y", "e", "r", "p"}
	Mergesort(arrayString)
	fmt.Println(arrayString)

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
