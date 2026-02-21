package main

import (
	"fmt"
	"time"
)

func sayHello(){
	fmt.Println("Hello from Goroutine")
	time.Sleep(2000 * time.Millisecond) //Sleep is used to simulate a long running task
	fmt.Println("sayHello function ended successfully")
}

func sayHii(){
	fmt.Println("Hi bro")
}

func main(){
	fmt.Println("Learning Goroutines")
	go sayHello()
	sayHii()

	time.Sleep(2000 * time.Millisecond)
}