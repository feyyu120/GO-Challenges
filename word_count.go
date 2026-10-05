package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	var sentence string
	fmt.Println("enter sentence")
	if scan.Scan() {
		sentence = scan.Text()
	}
	holder := strings.Fields(sentence)
	dict := map[string]int{}
	for _, value := range holder {
		if _, exist := dict[value]; exist {
			dict[value]++
		} else {
			dict[value] = 1
		}
	}

	for i, j := range dict {
		fmt.Println(i, j)
	}
}
