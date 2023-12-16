package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func processTask() {
	// Simulate processing by sleeping for 2 seconds
	time.Sleep(2 * time.Second)
}

func processSequential(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing sequentially...")
	processTask()
	fmt.Fprintf(w, "Sequential processing complete.")
}

func processConcurrent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing concurrently...")
	var wg sync.WaitGroup

	// Simulate concurrent processing by launching multiple goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			fmt.Printf("Worker %d processing...\n", workerID)
			processTask()
			fmt.Printf("Worker %d complete.\n", workerID)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Fprintf(w, "Concurrent processing complete.")
}

func main() {
	// Define handlers for the two endpoints
	http.HandleFunc("/process-single", processSequential)
	http.HandleFunc("/process-concurrent", processConcurrent)

	// Start the server and listen on port 8000
	fmt.Println("Server listening on :8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
