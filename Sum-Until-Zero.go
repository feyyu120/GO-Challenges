package main

import "fmt"

func main() {
	sum := 0
	var num int = 1
	fmt.Println("Enter the integer")
	for num != 0 {
		fmt.Scan(&num)
		sum += num
	}
	fmt.Print("The sum of the previous inputs is: ")
	fmt.Println(sum)
}
