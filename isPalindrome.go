package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var num int
	//create list
	lst := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		lst = append(lst, num)
	}
	//using two pointer to check it's palindrome
	left := 0
	right := len(lst) - 1
	count := 0
	for left < right {
		if lst[left] != lst[right] {
			fmt.Println("NO")
			break
		}
		count++
		left++
		right--
	}

	if count == n/2 {
		fmt.Println("YES")
	}

}
