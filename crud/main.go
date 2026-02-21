package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	// "io/ioutil"
	"net/http"
)

type Todo struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`

}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error getting GET response", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error getting response", res.StatusCode)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response", err)
	// 	return
	// }
	// fmt.Println("Data:", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo) //Decode takes a byte slice and a pointer to the struct where the data will be stored
	if err != nil{
		fmt.Println("Error decoding", err)
		return
	}
	fmt.Println("Todo:", todo)
}

func performPostRequest() {
	todo := Todo{
		UserId: 23,
		Title: "Harsh Raj",
		Completed: true,
	}

	//convert todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil{
		fmt.Println("Error Marshalling:",err)
		return
	}
	//convert jsonData to string and print
	jsonString := string(jsonData);

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"
	
	//send POST request with json data
	res, err := http.Post(myUrl, "application/json", jsonReader)
	
	if err != nil{
		fmt.Println("Error sending POST request", err)
		return
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println("Error in reading",err)
		return
	}
	fmt.Println("Response from POST request:", string(data))

	fmt.Println("Response status code:", res.StatusCode)
}

func performPutRequest() {
	todo := Todo{
		UserId: 66,
		Title: "Harsh Raj Learning GoLang",
		Completed: false,
	}

	//convert todo struct to json
	jsondData, err := json.Marshal(todo)
	if err != nil{
		fmt.Println("Error Marshalling:",err)
		return
	}

	//convert jsonData to string and print
	jsonString := string(jsondData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const url = "https://jsonplaceholder.typicode.com/todos/1"

	//create a new request with method PUT, url and json data
	req, err := http.NewRequest(http.MethodPut, url, jsonReader)
	if err != nil{
		fmt.Println("Error creating request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	//send the request using http.
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil{
		fmt.Println("Error sending request", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println("Error reading response", err)
		return
	}
	fmt.Println("Response from PUT request:", string(data))

	fmt.Println("Status code", res.StatusCode)
}

func performDeleteRequest() {
	const url = "https://jsonplaceholder.typicode.com/todos/1"

	//create a new request with method DELETE and url
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil{
		fmt.Println("Error creating in Delete request:",err)
		return
	}

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil{
		fmt.Println("Error sending request", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)
}

func main() {
	fmt.Println("Learning CRUD operations")
	// performGetRequest()	
	// performPostRequest()
	// performPutRequest()
	performDeleteRequest()
}