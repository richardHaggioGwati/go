package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Functions in Golang")

	result := superAdd(3, 7, 2, 9, 4)
	fmt.Println("The addition is:", result)

	roll, message := diceRoll()
	fmt.Println("Dice roll:", roll)
	fmt.Println("Message:", message)

	alice := person {
		name:        "Alice",
		age:		 30,
		ethnicity:   "Caucasian",
		dateOfBirth: "1993-05-15",
		hairColor:   "Brown",
		eyeColor:    "Blue",
	}

	fmt.Println(alice.greet())
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func diceRoll() (int, string) {
	roll := rand.IntN(6) + 1
	var message string

	switch roll {
	case 1:
		message = "You rolled a one. Try again!"
	case 2:
		message = "You rolled a two. Not bad!"
	case 3:
		message = "You rolled a three. Keep going!"
	case 4:
		message = "You rolled a four. Nice!"
	case 5:
		message = "You rolled a five. Great!"
	case 6:
		message = "You rolled a six! Great job!"
	default:
		message = "You rolled a decent number."
	}
	return roll, message
}

type person struct {
	name        string
	age         int
	ethnicity   string
	dateOfBirth string
	hairColor   string
	eyeColor    string
}

func (p person) greet() string {
	return fmt.Sprintf("Hello, my name is %s. I am %d years old.", p.name, p.age)
}
