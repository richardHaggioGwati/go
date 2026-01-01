package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	// Creating a map
	var colors = map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println("Colors:", colors)

	// Accessing a map
	fmt.Println("Color red code:", colors["red"])
	
	// Adding a new key-value pair
	colors["yellow"] = "#ffff00"
	fmt.Println("Colors after adding yellow:", colors)

	// Modifying a value
	colors["green"] = "#008000"
	fmt.Println("Colors after modifying green:", colors)

	// Deleting a key-value pair
	delete(colors, "blue")
	fmt.Println("Colors after deleting blue:", colors)

	// Looping through a map

	for color, hex := range colors {
		fmt.Printf("Color: %s, Hex: %s\n", color, hex)
	}
}