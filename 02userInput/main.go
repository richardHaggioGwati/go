package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("User input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a random value: ")

	// comma ok pattern

	input, _ := reader.ReadString('\n')

	fmt.Printf("You entered: %s", input)
	fmt.Printf("Value is of type %T \n", input)
	fmt.Println("End of program")
}
	