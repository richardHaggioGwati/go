package main

import (
	"strings"
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"password,omitempty"`
	Tags     []string `json:"tags"`
}

func main() {
	// Encode JSON data
	courseData := []course{
		{
			Name:     "Learn Golang",
			Price:    299,
			Platform: "learncodeonline.in",
			Password: "abcd123",
			Tags:     []string{"programming", "golang", "backend"},
		},
		{
			Name:     "Learn Python",
			Price:    199,
			Platform: "learncodeonline.in",
			Password: "xyz987",
			Tags:     []string{"programming", "python", "data science"},
		},
		{
			Name:     "Learn JavaScript",
			Price:    149,
			Platform: "learncodeonline.in",
			Password: "js456",
			Tags:     []string{"programming", "javascript", "frontend"},
		},
	}

	encodedJson := encodeJson(courseData)
	// decodeJson := decodeJson(encodedJson)
	unknownFieldsJson, _ := decodeJsonWithUnknownFields(encodedJson)
	//fmt.Printf("Decoded JSON: %#v\n", decodeJson)
	// fmt.Println(string(encodedJson))
	fmt.Println("JSON with unknown fields:\n", unknownFieldsJson)
}

func encodeJson(courses []course) []byte {
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	return finalJson
}

func decodeJson(jsonData []byte) []course {
	var course []course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &course)
		fmt.Printf("%#v\n", course)
	} else {
		fmt.Println("JSON is not valid")
	}
	return course
}

func decodeJsonWithUnknownFields(jsonData []byte) (string, error) {
	var data []map[string]interface{}

	if err := json.Unmarshal(jsonData, &data); err != nil {
		return "", err
	}

	var result strings.Builder
	for i, item := range data {
		result.WriteString(fmt.Sprintf("Course #%d\n", i+1))
		for key, value := range item {
			result.WriteString(fmt.Sprintf("  Key: %s, Value: %#v\n", key, value))
		}
	}

	return result.String(), nil
}