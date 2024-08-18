# Abstract Parallel Processing in Go

## Simplifies Concurrent Processing
**Abstracts Complexity**: ParallelForEach makes it easy to work with concurrency by taking care of goroutines, synchronizing with sync.WaitGroup, and handling channels. Developers can focus on buisness logic without worrying about the details of how concurrency works.

**Easy to Use**: Developers only need to specify the list of items, the maximum number of concurrent goroutines, and the function to apply to each item. The rest is handled by ParallelForEach.

The project includes a custom `ParallelForEach` function that processes a list of items in parallel, using a configurable number of goroutines. The processing function includes a simulated delay to mimic time-consuming tasks.

## How It Works

The `ParallelForEach` function takes a list of items and processes each item concurrently using a specified number of goroutines. The number of goroutines can be adjusted to balance performance and resource usage.

```
maxGoroutines := 10
// Call ParallelForEach from the parallel package to process the objects and collect results.
results := parallel.ParallelForEach(items, maxGoroutines, ProcessObject)
```
The processing function, `ProcessObject`, simulates a time-consuming task by including a 1-second delay.

## Scenarios

### Scenario 1: Number of Items = 10, maxGoroutines = 1

- **Description:** In this scenario, we process 10 items with only 1 goroutine. Since there is only one goroutine, the items will be processed sequentially.
- **Expected Behavior:** The total processing time will be approximately equal to the number of items (approx 10 seconds), as each item is processed one after another.
- **Expected Output:**

  ```
  Result from processing item 7: 14
  Result from processing item 8: 16
  Result from processing item 9: 18
  Result from processing item 10: 20
  Time taken: 10.0965608s
  ```


### Scenario 2: Number of Items = 10, maxGoroutines = 10

- **Description:** In this scenario, we process 10 items with 10 goroutines. Since each item has its own goroutine, all items are processed concurrently.
- **Expected Behavior:** The total processing time will be approximately equal to the time taken by a single item (approx 1 second), as all items are processed simultaneously.

- **Expected Output:**


  ```
    Result from processing item 7: 14
    Result from processing item 8: 16
    Result from processing item 9: 18
    Result from processing item 10: 20
    Time taken: 1.0136141s
  ```
