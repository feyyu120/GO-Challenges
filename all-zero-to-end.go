package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter the size of the array")
	fmt.Scan(&n)
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			nums = append(nums, num)
		} else {
			nums = append([]int{num}, nums...)
		}
	}
	fmt.Println(nums)
}
