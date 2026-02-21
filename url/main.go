package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"
	fmt.Printf("Type of url:%T\n", myUrl)

	parsedUrl, err := url.Parse(myUrl)
	if err != nil{
		fmt.Println("Can't parse URL", err)
		return
	}
	fmt.Printf("Type of URL: %T\n", parsedUrl)

	fmt.Println("Scheme: ", parsedUrl.Scheme) //Scheme is the protocol used in the URL, such as "http" or "https"
	fmt.Println("Host: ", parsedUrl.Host) //Host is the domain name or IP address of the server, along with an optional port number
	fmt.Println("Path: ", parsedUrl.Path) //Path is the specific resource on the server that the URL points to, such as "/todos/1"
	fmt.Println("RawQuery: ", parsedUrl.RawQuery) //RawQuery is the query string of the URL, which contains any parameters that are passed to the server, such as "id=1&name=John"

	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "username=JohnDoe"

	//create a new URL string from the parsed URL
	newUrl := parsedUrl.String()
	fmt.Println("New URL:", newUrl)


}
