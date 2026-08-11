package main

import "fmt"

type Courses struct {
	Name        string
	Point       float64
	Credit_hour int
}

func main() {
	var course []Courses
	var n int
	fmt.Println("enter the the length of the course you take")
	fmt.Scan(&n)
	fmt.Println("Enter name and credit hour of the course")
	for i := 0; i < n; i++ {
		var name string
		var point float64
		var credit_hour int
		fmt.Scan(&name, &point, &credit_hour)
		course = append(course, Courses{Name: name, Credit_hour: credit_hour})
	}

	fmt.Println(course[0].Credit_hour)
}

func Calculate(name string, credit, n int) {

}
