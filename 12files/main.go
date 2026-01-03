package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in Golang")
	content := "Hello, World! is a common first program in programming."

	file, err := os.Create("hello.txt")
	checkErrNil(err)

	length, err := io.WriteString(file, content)
	checkErrNil(err)

	fmt.Println("Length of file content is:", length)
	defer file.Close()

	readFile("./hello.txt")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	checkErrNil(err)

	fmt.Println("File content is:", string(data))

}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}