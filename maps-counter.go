package main

import "fmt"

func countWords(words []string) map[string]int {
	result := make(map[string]int)
	for _, j := range words {
		result[j] += 1
	}
	return result
}

func main() {
	input := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	fmt.Println(countWords(input))
}
