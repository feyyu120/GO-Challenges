package main

import "fmt"

func MaxSub(lst []int) int {
	current := lst[0]
	bestSum := lst[0]
	for i := 1; i < len(lst); i++ {
		if current > lst[i]+current {
			current = lst[i]
		} else {
			current += lst[i]
			if bestSum < current {
				bestSum = current
			}
		}
	}

	return bestSum
}

func main() {
	lst := []int{5, -2, 3, -1, 2}
	result := MaxSub(lst)
	fmt.Println(result)
}
