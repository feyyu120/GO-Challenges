package main

import (
	"fmt"
)

func MaxSum(lst []int, k int) int {
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += lst[i]
	}
	maxSum := 0
	for right := k; right < len(lst); right++ {
		windowSum += lst[right]
		windowSum -= lst[right-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

func main() {
	lst := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	k := 4
	result := MaxSum(lst, k)
	fmt.Println(result)
}
