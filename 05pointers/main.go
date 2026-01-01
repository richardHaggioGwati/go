package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	myNumber := 27

	var pointer *int = &myNumber

	fmt.Println("Value of myNumber:", myNumber)
	fmt.Println("Address of myNumber:", &myNumber)
	fmt.Println("Value of pointer:", pointer)
	fmt.Println("Value at the address stored in pointer:", *pointer + 5)
}