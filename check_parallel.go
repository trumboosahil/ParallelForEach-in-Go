package main

import (
	"concurrency_parallel/parallel" // Import the custom parallel package
	"fmt"
	"time"
)

// Object represents the type of objects you want to process.
type Object struct {
	ID    int
	Value int
}

// Mock the user defined function and introduce delay of 1 second
func ProcessObject(obj Object) int {
	time.Sleep(1 * time.Second)
	return obj.Value * 2
}

func main() {
	// Create a list of ite
	timeStart := time.Now()
	items := make([]Object, 10)
	for i := 0; i < 10; i++ {
		items[i] = Object{
			ID:    i + 1,
			Value: i + 1,
		}
	}

	maxGoroutines := 10

	// Call ParallelForEach from the parallel package to process the objects and collect results.
	results := parallel.ParallelForEach(items, maxGoroutines, ProcessObject)

	// Print the results.
	for i, result := range results {
		fmt.Printf("Result from processing item %d: %d\n", items[i].ID, result)
	}

	fmt.Printf("Time taken: %s\n", time.Since(timeStart))
}
