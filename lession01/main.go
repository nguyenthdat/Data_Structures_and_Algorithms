// Lesson 1 - Binary Search, Linked Lists and Complexity
package main

import (
	"fmt"
	"sort"
)

func main() {
	cards := []int{13, 11, 12, 7, 4, 3, 1, 0}
	fmt.Println(locateCard(cards, 7))
}

func locateCard(cards []int, query int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(cards))) // To sort in descending order, we use sort.Reverse function.
	for i, v := range cards {
		if v == query {
			return i
		}
	}
	return 0
}
