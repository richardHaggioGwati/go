package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=123"

func main() {
	fmt.Println("Handling urls in Golang")

	//parse url

	parsedUrl, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Raw Query:", parsedUrl.RawQuery)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Hostname:", parsedUrl.Hostname())

	qParams := parsedUrl.Query()
	fmt.Println("Query Params:", qParams)
	fmt.Println("Course Name:", qParams["coursename"])
	fmt.Println("Payment ID:", qParams["paymentid"])

	for _, val := range qParams {
		fmt.Println("Param value:", val)
	}

	//construct url
	anotherUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	fmt.Println("Another URL is:", anotherUrl.String())
}