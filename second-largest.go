package main

import "fmt"

func main() {

	var n int
	fmt.Println("enter the length of the numbers")
	fmt.Scan(&n)
	var firstlarge int = 0
	var secondlarge int = 0
	fmt.Println("enters numbers")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		if firstlarge < num {
			secondlarge = firstlarge
			firstlarge = num

		}
		if i == 0 {
			secondlarge = num
		}

	}
	fmt.Println(secondlarge)
}
