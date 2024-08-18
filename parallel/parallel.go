// parallel/parallel.go
package parallel

import (
	"sync"
)

// ParallelForEach processes a list of items in parallel and collects the results.
func ParallelForEach[T any, R any](items []T, maxGoroutines int, fn func(T) R) []R {
	var wg sync.WaitGroup
	guard := make(chan struct{}, maxGoroutines)

	// Create a slice to store the results.
	results := make([]R, len(items))

	for i, item := range items {
		wg.Add(1)
		guard <- struct{}{}

		go func(i int, item T) {
			defer wg.Done()
			defer func() { <-guard }()
			// Call the user-defined processing function and store the result.
			results[i] = fn(item)
		}(i, item)
	}

	wg.Wait() // Wait for all goroutines to finish.

	return results
}
