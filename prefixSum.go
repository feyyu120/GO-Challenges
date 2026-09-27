package main

import (
	"fmt"
)

func Sum(lst []int, a, b int) int {
	sum := 0

	for _, value := range lst[a : b+1] {
		sum += value
	}
	return sum
}

func main() {

	lst := []int{3, 2, 4, 1, 5}
	// ranges
	var (
		a, b int
	)
	fmt.Scan(&a, &b)
	result := Sum(lst, a, b)
	fmt.Println(result)
}
