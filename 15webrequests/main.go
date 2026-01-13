package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Get request in Golang")
	const baseURL = "http://localhost:8000"

	// Perform a GET request
	getRequest(baseURL + "/get")

	// Perform a POST request with JSON payload
	jsonPayload := strings.NewReader(`{
		"courseName": "Learn Golang",
		"price": 299,
		"platform": "learncodeonline.in"
	}`)
	postJsonRequest(baseURL+"/post", jsonPayload)

//	Perform a POST request with form data
	formData := url.Values{}
	formData.Add("firstName", "John")
	formData.Add("lastName", "Doe")
	formData.Add("email", "john.doe@example.com")
	postFormRequest(baseURL+"/postform", formData)
}

func getRequest(url string) {
	response, err := http.Get(url)
	errorCheck(err)

	defer response.Body.Close()
	fmt.Println("Response status code:", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)

	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)
	errorCheck(err)

	// Alternative way to read the response body using a string builder
	responseString.Write(content)
	fmt.Println("Response body is:", responseString.String())

	//fmt.Println("Response body is:", string(content))
}

func postJsonRequest(url string, jsonPayload *strings.Reader) {
	response, err := http.Post(url, "application/json", jsonPayload)
	errorCheck(err)
	defer response.Body.Close()

	fmt.Println("Response status code:", response.StatusCode)
	content, err := io.ReadAll(response.Body)
	errorCheck(err)

	fmt.Println("Response body is:", string(content))
}

func postFormRequest(link string, data url.Values) {
	response, err := http.PostForm(link, data)
	errorCheck(err)
	defer response.Body.Close()

	fmt.Println("Response status code:", response.StatusCode)
	content, err := io.ReadAll(response.Body)

	errorCheck(err)
	fmt.Println("Response body is:", string(content))

}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}