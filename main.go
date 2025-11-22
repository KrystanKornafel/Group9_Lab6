package main

import (
	"fmt"
	"time"
)

//"sync"
//"time"

//to run, simply say in the terminal "go run practiceTasks.go" in the terminal
//code written in VS code

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(500 * time.Millisecond)
		results <- j * 2
	}
}
func main() {
	//4th task
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	//Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	//Collect results
	for i := 1; i <= 5; i++ {
		fmt.Println("Results:", <-results)
	}

}
