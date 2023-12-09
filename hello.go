package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

// Request structure for receiving arrays in the request body
type SortRequest struct {
	ToSort [][]int `json:"to_sort"`
}

// Response structure for returning the sorted arrays and time taken
type SortResponse struct {
	SortedArrays [][]int       `json:"sorted_arrays"`
	TimeNS       int64         `json:"time_ns"`
}

func main() {
	// Define the HTTP server and endpoints
	http.HandleFunc("/process-single", processSingle)
	http.HandleFunc("/process-concurrent", processConcurrent)

	// Start the server on port 8000
	fmt.Println("Server is listening on port 8000...")
	http.ListenAndServe(":8000", nil)
}

func processSingle(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON request
	var request SortRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Measure the time taken for sequential sorting
	startTime := time.Now()
	sortedArrays := sortSequential(request.ToSort)
	timeTaken := time.Since(startTime).Nanoseconds()

	// Return the sorted arrays and time taken in nanoseconds in the response
	response := SortResponse{
		SortedArrays: sortedArrays,
		TimeNS:       timeTaken,
	}
	sendJSONResponse(w, response)
}

func processConcurrent(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON request
	var request SortRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Measure the time taken for concurrent sorting
	startTime := time.Now()
	sortedArrays := sortConcurrent(request.ToSort)
	timeTaken := time.Since(startTime).Nanoseconds()

	// Return the sorted arrays and time taken in nanoseconds in the response
	response := SortResponse{
		SortedArrays: sortedArrays,
		TimeNS:       timeTaken,
	}
	sendJSONResponse(w, response)
}

func sortSequential(input [][]int) [][]int {
	result := make([][]int, len(input))
	for i, arr := range input {
		result[i] = make([]int, len(arr))
		copy(result[i], arr)
		sort.Ints(result[i])
	}
	return result
}

func sortConcurrent(input [][]int) [][]int {
	var wg sync.WaitGroup
	wg.Add(len(input))

	result := make([][]int, len(input))

	for i, arr := range input {
		go func(i int, arr []int) {
			defer wg.Done()
			result[i] = make([]int, len(arr))
			copy(result[i], arr)
			sort.Ints(result[i])
		}(i, arr)
	}

	wg.Wait()
	return result
}

func sendJSONResponse(w http.ResponseWriter, response SortResponse) {
	// Set response headers and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send the JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

/*
# Sequential sorting
curl -X POST -H "Content-Type: application/json" -d '{"to_sort":[[3,2,1],[6,5,4],[9,8,7]]}' http://localhost:8000/process-single

# Concurrent sorting
curl -X POST -H "Content-Type: application/json" -d '{"to_sort":[[3,2,1],[6,5,4],[9,8,7]]}' http://localhost:8000/process-concurrent
*/
