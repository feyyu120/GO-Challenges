package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var nums []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	maxVal := nums[0] + nums[1]
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			nums[i] + nums[j]
		}
	}
}
