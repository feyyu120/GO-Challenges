package main

import (
	"fmt"
	"math/big"
)

func factorial(n int64) *big.Int {
	// Initialize result as a big.Int with a starting value of 1
	result := big.NewInt(1)

	// Loop up to n
	for i := int64(1); i <= n; i++ {
		// Convert the loop index 'i' into a temporary big.Int
		bigI := big.NewInt(i)

		// Multiply: result = result * bigI
		result.Mul(result, bigI)
	}

	return result
}

func main() {
	var num int64
	fmt.Println("Enter the number:")
	fmt.Scan(&num)

	result := factorial(num)

	// %s prints the entire, giant number as a clean text string
	fmt.Printf("The factorial of %d is:\n%d\n", num, result)
}
