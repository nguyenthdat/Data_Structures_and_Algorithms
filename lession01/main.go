// Lesson 1 - Binary Search, Linked Lists and Complexity
package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{13, 11, 12, 7, 4, 3, 1, 0}
	fmt.Println(locateCard(slice, 7))
}

func locateCard(slice []int, index int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(slice))) // To sort in descending order, we use sort.Reverse function.
	for i, v := range slice {
		if v == index {
			return i
		}
	}
	return 0
}
