package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"` //json tag is used to specify the name of the field in the JSON object
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Learning JSON")
	person := Person{Name: "John Doe", Age: 30, IsAdult: true}
	// fmt.Println("Person data:", person)

	//convert person struct to JSON Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil{
		fmt.Println("Error marshalling",err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

	//Decoding (Unmarshalling) JSON data back to struct
	var personData Person
	err = json.Unmarshal(jsonData, &personData) //Unmarshal takes a byte slice and a pointer to the struct where the data will be stored
	if err != nil{
		fmt.Println("Error unmarshalling", err)
		return
	}
	fmt.Println("Person data after unmarshalling:", personData)
}
