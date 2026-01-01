package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go")

	// Creating a slice
	var veggies = []string{"carrot", "potato", "cabbage"}
	fmt.Println("Veggies:", veggies)

	// Slicing a slice
	var someVeggies = veggies[1:3]
	fmt.Println("Some Veggies:", someVeggies)

	// Modifying a slice
	someVeggies[0] = "beans"
	fmt.Println("Modified Some Veggies:", someVeggies)
	fmt.Println("Original Veggies after modification:", veggies)

	// Appending to a slice
	someVeggies = append(someVeggies, "spinach")
	fmt.Println("Appended Some Veggies:", someVeggies)
	fmt.Println("Original Veggies after appending:", veggies)

	// Length and Capacity
	fmt.Println("Length of Some Veggies:", len(someVeggies))
	fmt.Println("Capacity of Some Veggies:", cap(someVeggies))

	// Creating a slice with make
	highScores := make([]int, 3)

	highScores[0] = 100
	highScores[1] = 200
	highScores[2] = 300
	// highScores[3] = 400  This will cause a runtime error: index out of range

	fmt.Println("High Scores:", highScores)

	highScores = append(highScores, 400, 500, 600)
	fmt.Println("High Scores after appending:", highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted High Scores:", highScores)

	// Removing an element from a slice

	var courses = []string{"reactjs", "javascript", "html", "css", "vuejs"}
	fmt.Println("Courses before removal:", courses)
	var index = 3
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses removed by index...", courses)
}
