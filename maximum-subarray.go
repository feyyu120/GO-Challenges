package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// Handle empty input edge case
	if n <= 0 {
		fmt.Println(0)
		return
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	// Initialize Kadane's Algorithm variables
	maxSoFar := nums[0]
	currentMax := nums[0]

	// Loop through the array starting from the second element
	for i := 1; i < n; i++ {
		// Decide whether to add the current number to the existing subarray
		// or start a brand new subarray from the current number
		if nums[i] > currentMax+nums[i] {
			currentMax = nums[i]
		} else {
			currentMax = currentMax + nums[i]
		}

		// Update the overall maximum found so far
		if currentMax > maxSoFar {
			maxSoFar = currentMax
		}
	}

	fmt.Println(maxSoFar)
}
