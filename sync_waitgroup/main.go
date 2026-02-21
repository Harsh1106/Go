package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup){
	defer wg.Done() //Decrement the WaitGroup counter when the worker is done
	fmt.Printf("Worker %d started\n",i)
	//some task is happening here
	fmt.Printf("Worker %d end\n",i)
}

func main(){
	//fmt.Println("Explore goroutine started")
	var wg sync.WaitGroup
	//start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) //Increment the WaitGroup counter for each worker
		go worker(i, &wg)
	}
	wg.Wait() //Wait for all workers to finish
	fmt.Println("worker task completed")
}