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

	//3rd task
	// counter := 0
	// ch := make(chan int)

	// // Start 1000 goroutines
	// for i := 0; i < 1000; i++ {
	// 	go func() { ch <- 1 }()
	// }

	// //Collect results
	// for i := 0; i < 1000; i++ {
	// 	counter += <-ch
	// }

	// fmt.Println("Final counter:", counter)

	//2nd task
	// var wg sync.WaitGroup
	// counter := 0

	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	//task 2
	// 	go func() {
	// 		defer wg.Done()
	// 		var mu sync.Mutex
	// 		mu.Lock()
	// 		counter++
	// 		mu.Unlock()
	// 	}()

	// }

	// wg.Wait()
	// fmt.Println("Final counter:", counter)

	// //1st exercise of the lab on goroutines
	// go sayHello("Goroutine 1")
	// go sayHello("Goroutine 2")
	// go sayHello("Goroutine 3")

	// time.Sleep(1 * time.Second) //crude wait so goroutines can run
}

// from task 1
// func sayHello(name string) {
// 	fmt.Println("Hello from", name)
// }
