package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, 0, n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	for _, j := range nums {
		isContain := slices.Contains(nums, j)
		var count int = 0
		for isContain {
			count++
			if ind := slices.Index(nums, j); ind != -1 {
				nums = slices.Delete(nums, ind, ind+1)
			}

			isContain = slices.Contains(nums, j)
		}
		fmt.Println(j, count)
	}
}
