package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Requests in Golang")

	response, error := http.Get(url)
	
	if error != nil {
		panic(error)
	}
	
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // callers responsibility to close the body

	databytes, error := ioutil.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}

	content := string(databytes)
	fmt.Println("Response content is:", content)
}