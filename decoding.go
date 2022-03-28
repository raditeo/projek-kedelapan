package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	// var jsonString = `
	// 	{
	// 		"full_name": "Raditeo Warma",
	// 		"email": "raditeo@walisongo.ac.id",
	// 		"age": 23
	// 	}
	// `

	var jsonString = `[
			{
				"full_name": "Raditeo Warma",
				"email": "raditeo@walisongo.ac.id",
				"age": 23
			},
			{
				"full_name": "Nadia Rizqi",
				"email": "nadia@walisongo.ac.id",
				"age": 21
			}
		]
	`

	// var result Employee
	// var result map[string]interface{}

	// var err = json.Unmarshal([]byte(jsonString), &result)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// var temp interface{}

	// var err = json.Unmarshal([]byte(jsonString), &temp)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// var result = temp.(map[string]interface{})

	// fmt.Println("full_name:", result.FullName)
	// fmt.Println("email:", result.Email)
	// fmt.Println("age:", result.Age)

	// fmt.Println("full_name:", result["full_name"])
	// fmt.Println("email:", result["email"])
	// fmt.Println("age:", result["age"])

	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("index %d: %+v\n", i+1, v)
	}
}
