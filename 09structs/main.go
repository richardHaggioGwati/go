package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	// Creating a User struct instance
	alice := User{
		Name:   "Alice",
		Email:  "alice@example.com",
		Age:    30,
		Status: true,
	}

	fmt.Printf("Alice's details are: %+v \n", alice)
	fmt.Printf("Name is %v, and the age is %v.", alice.Name, alice.Age)

}

/*
User struct represents a user with basic information.
*/
type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
