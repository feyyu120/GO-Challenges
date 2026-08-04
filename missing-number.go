package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter size the integers")
	fmt.Scan(&n)
	var missed int
	for i := 1; i < n; i++ {
		var num int
		fmt.Scan(&num)
		if i != num {
			missed = i
		}
	}
	fmt.Println(missed)

}
