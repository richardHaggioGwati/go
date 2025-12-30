package main

import "fmt"

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
}