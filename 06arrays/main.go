package main

import "fmt"

func main() {

	fmt.Println("Arrays")

	var grades [3]int

	grades[0] = 97
	grades[2] = 93

	fmt.Println("Grades:", grades)
	fmt.Println("Number of grades:", len(grades))

	var students = [3]string{"Alice", "Bob", "Charlie"}

	fmt.Println("Students:", students)
	fmt.Println("Number of students:", len(students))

}