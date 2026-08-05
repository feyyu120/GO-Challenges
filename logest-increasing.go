package main

import "fmt"

func main() {
	var count int = 1
	var prev int = 0
	var maxs []int
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num-prev == 1 && num != 1 {
			count++
		}
		if num-prev != 1 {
			maxs = append(maxs, count)
			count = 1
		}
		prev = num
	}
	maxs = append(maxs, count)
	maxVal := maxs[0]
	for i := 0; i < len(maxs); i++ {
		if maxs[i] > maxVal {
			maxVal = maxs[i]
		}
	}
	fmt.Print(maxVal)
}
