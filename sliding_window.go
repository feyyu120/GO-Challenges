package main

import (
	"fmt"
)

func main() {
	lst := []int{2, 1, 5, 1, 3, 2}

	maxSum := 0
	windowSum := 0
	k := 3
	for j := 0; j < k; j++ {
		windowSum += lst[j]
	}
	maxSum = windowSum
	for right := k; right < len(lst); right++ {

		windowSum += lst[right]
		windowSum -= lst[right-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	fmt.Println(maxSum)
}
