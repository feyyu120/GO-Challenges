package main

import "fmt"

type Courses struct {
	Name        string
	Point       float64
	Credit_hour int
}

func main() {
	var course []Courses
	var total int = 0
	var n int
	fmt.Println("enter the the length of the course you take")
	fmt.Scan(&n)
	fmt.Println("Enter name,point and credit hour of the course")
	for i := 0; i < n; i++ {
		var name string
		var point float64
		var credit_hour int
		fmt.Scan(&name, &point, &credit_hour)
		total += credit_hour
		course = append(course, Courses{Name: name, Point: point, Credit_hour: credit_hour})
	}
	Calculate(course, n, total)

}

func Calculate(course []Courses, n int, total int) {
	result := 0
	for i := 0; i < n; i++ {
		c := course[i].Credit_hour
		val := course[i].Point
		switch {
		case val > 85:
			result += 4 * c
		case val > 70:
			result += 3 * c
		}
	}
	if total != 0 {
		fmt.Println("Your grade is", result/total)
	}

}
