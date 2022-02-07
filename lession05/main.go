package main

import "fmt"

func main() {
	fmt.Println(getFactorial(5))
}

func getFactorial(i int) int {
	if i < 2 {
		return 1
	}
	return i * getFactorial(i-1)
}
