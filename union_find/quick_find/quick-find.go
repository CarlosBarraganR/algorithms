package main

import "fmt"

/*
 * Quick Find (complex)
 * Complexity: O2 [cuadratic]
 */
func main() {
	array := make([]int, 10)

	for i := 0; i < 10; i++ {
		array[i] = i
	}

	fmt.Println(array) // [0 1 2 3 4 5 6 7 8 9]

	fmt.Println(connected(0, 9, &array)) // false
	union(0, 9, &array)
	fmt.Println(array)                   // [9 1 2 3 4 5 6 7 8 9]
	fmt.Println(connected(0, 9, &array)) // true
	union(1, 9, &array)

	fmt.Println(array) // [9 9 2 3 4 5 6 7 8 9]
}

func union(p int, q int, array *[]int) {
	for i := 0; i < len(*array); i++ {
		if (*array)[i] == (*array)[p] {
			(*array)[i] = (*array)[q]
		}
	}
}

func connected(p int, q int, array *[]int) bool {
	return (*array)[p] == (*array)[q]
}
