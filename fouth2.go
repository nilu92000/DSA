package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"
)

type Payload struct {
	ToSort [][]int `json:"to_sort"`
}

type SortedResponse struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNs       int64   `json:"time_ns"`
}

func processSingle(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	// Decode JSON payload
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Sequential sorting
	sortedArrays := sequentialSort(payload.ToSort)

	// Calculate the time taken
	timeTaken := time.Since(startTime).Nanoseconds()

	// Prepare response
	response := SortedResponse{
		SortedArrays: sortedArrays,
		TimeNs:       timeTaken,
	}

	// Encode and send the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sequentialSort(toSort [][]int) [][]int {
	var sortedArrays [][]int

	for _, innerList := range toSort {
		sortedInnerList := make([]int, len(innerList))
		copy(sortedInnerList, innerList)
		sort.Ints(sortedInnerList)
		sortedArrays = append(sortedArrays, sortedInnerList)
	}

	return sortedArrays
}

func main() {
	// Define endpoints
	http.HandleFunc("/process-single", processSingle)

	// Start server
	port := ":8000"
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
