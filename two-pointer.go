package main

import (
	"fmt"
)

func main() {
	lst := []int{1, 2, 3, 4, 5}
	left := 0
	right := len(lst) - 1

	for left < right {
		lst[left], lst[right] = lst[right], lst[left]
		left++
		right--
	}
	fmt.Println(lst)
}
