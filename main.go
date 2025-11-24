package main

import (
	"fmt"
	"io"
	"net/http"
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
	for j := range jobs {
		//use jobs?
		//hint: use resp, err := http.Get(url)
		resp, err := http.Get(j)
		//TODO: send Result struct to results channel
		//if unsuccessful, do this
		if err != nil {
			results <- FetchResult{URL: j, Error: err}
			//continue breaks out of this error, and moves to the top of the loop again to go to the next job in the list
			continue
		}
		//successful case
		body, err := io.ReadAll(resp.Body)
		results <- FetchResult{URL: j, StatusCode: resp.StatusCode, Size: len(body), Error: err}
		defer resp.Body.Close()
	}

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
		jobs <- urls[j-1]
	}
	close(jobs)

	//TODO: Collect results
	for i := 1; i <= numJobs; i++ {
		fmt.Println("Results:", <-results)
	}

	fmt.Println("\nScraping complete!")
}
