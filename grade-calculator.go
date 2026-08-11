package main

import "fmt"

type courses struct {
	Name        string
	Credit_hour int
}

func main() {
	var n int
	course := make([]courses, 0, n)
	fmt.Println("enter the the length of the course you take")
	fmt.Scan(&n)
	fmt.Println("Enter name and credit hour of the course")
	for i := 0; i < n; i++ {
		var name string
		var credit_hour int
		fmt.Scan(&name, &credit_hour)
		course[i].Name = name
		course[i].Credit_hour = credit_hour
	}

	fmt.Println(course)
}
