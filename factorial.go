package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	return n * factorial(n-1)

}

func main() {
	var num int
	fmt.Println("Enter the number:")
	fmt.Scan(&num)

	result := factorial(num)

	// %s prints the entire, giant number as a clean text string
	fmt.Printf("The factorial of %d is:\n%d\n", num, result)
}
