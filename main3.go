package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"sync"
)

type Payload struct {
	ToSort [][]int `json:"to_sort"`
}

type SortedResponse struct {
	SortedLists [][]int `json:"sorted_lists"`
}

func processSingle(w http.ResponseWriter, r *http.Request) {
	process(w, r, false)
}

func processConcurrent(w http.ResponseWriter, r *http.Request) {
	process(w, r, true)
}

func process(w http.ResponseWriter, r *http.Request, concurrent bool) {
	// Decode JSON payload
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Sequential or concurrent sorting
	var sortedLists [][]int
	if concurrent {
		sortedLists = concurrentSort(payload.ToSort)
	} else {
		sortedLists = sequentialSort(payload.ToSort)
	}

	// Prepare response
	response := SortedResponse{SortedLists: sortedLists}

	// Encode and send the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sequentialSort(toSort [][]int) [][]int {
	var sortedLists [][]int

	for _, innerList := range toSort {
		sortedInnerList := make([]int, len(innerList))
		copy(sortedInnerList, innerList)
		sort.Ints(sortedInnerList)
		sortedLists = append(sortedLists, sortedInnerList)
	}

	return sortedLists
}

func concurrentSort(toSort [][]int) [][]int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var sortedLists [][]int

	for _, innerList := range toSort {
		wg.Add(1)
		go func(innerList []int) {
			defer wg.Done()

			sortedInnerList := make([]int, len(innerList))
			copy(sortedInnerList, innerList)
			sort.Ints(sortedInnerList)

			mu.Lock()
			sortedLists = append(sortedLists, sortedInnerList)
			mu.Unlock()
		}(innerList)
	}

	wg.Wait()

	return sortedLists
}

func main() {
	// Define endpoints
	http.HandleFunc("/process-single", processSingle)
	http.HandleFunc("/process-concurrent", processConcurrent)

	// Start server
	port := ":8000"
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
