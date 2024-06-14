package main

import (
	"fmt"
	"time"
)

type Array struct {
	items []int
}

func (array *Array) sort() {
	for i := range array.items {
		for j := i + 1; j < len(array.items); j++ {
			if array.items[j] < array.items[i] {
				array.swap(i, j)
			}
		}
	}
}

func (array *Array) swap(a int, b int) {
	oldValue := array.items[a]
	array.items[a] = array.items[b]
	array.items[b] = oldValue
}

func main() {
	start := time.Now()
	array := &Array{[]int{14, 6, 232, 5, 63, 2, 1, 5}}
	fmt.Println(array) // [14 6 232 5 63 2 1 5]
	array.sort()
	fmt.Println(array) // [1 2 5 5 6 14 63 232]

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
