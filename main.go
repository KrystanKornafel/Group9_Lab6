package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// to run, simply say in the terminal "go run practiceTasks.go" in the terminal
// code written in VS code
type FetchResult struct {
	URL        string
	StatusCode int
	Size       int
	Error      error
}

func worker(id int, jobs <-chan string, results chan<- FetchResult) {
	//defer wg.Done()
	//TODO: fetch the URL
	//To read the URL from the terminal, you can do this
	//Reference used: https://www.geeksforgeeks.org/go-language/golang-program-to-check-status-of-a-url-website/
	// var url string
	// fmt.Print("Enter the URL of the website: ")
	// fmt.Scan(&url)
	for j := 1; j <= len(jobs); j++ {
		if j == id {
			var url string
			fmt.Print("Enter the URL of the website: ")
			fmt.Scan(&url)
			//use jobs?
			resp, err := http.Get(url)
			//TODO: send Result struct to results channel
			if err != nil {
				results <- FetchResult{URL: url, Error: err}
			}
			defer resp.Body.Close()
		}
	}
	//results <- res
	//hint: use resp, err := http.Get(url)
}
func main() {
	//var wg sync.WaitGroup
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://uottawa.ca",
		"https://github.com",
		"https://httpbin.org/get",
	}

	numWorkers := 5
	numJobs := 5

	jobs := make(chan string, len(urls))
	results := make(chan FetchResult, len(urls))

	//TODO: Start workers
	for w := 1; w <= numWorkers; w++ {
		//wg.Add(1)
		go worker(w, jobs, results)
	}

	//TODO: Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- strconv.Itoa(j)
	}
	close(jobs)

	//TODO: Collect results
	for i := 1; i <= numJobs; i++ {
		fmt.Println("Results:", <-results)
	}

	fmt.Println("\nScraping complete!")
}
