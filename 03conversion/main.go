package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("Conversions")
	fmt.Println("Enter a integer: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	userNumber, error := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if error != nil {
		fmt.Println("Error reading input:", error)
	} else {
		fmt.Println("You entered: ", userNumber)
		fmt.Printf("Value is of type %T \n", userNumber)
		fmt.Println("End of program")
	}
}