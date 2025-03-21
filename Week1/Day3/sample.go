package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printNums(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond) // Simulate delay
	}
}

func main() {
	start := time.Now()
	numOfWorkers := runtime.NumCPU()
	n := 1000
	chunkSize := n / numOfWorkers
	var wg sync.WaitGroup

	fmt.Printf("Using %d workers to print %d numbers\n", numOfWorkers, n)

	// Distribute the numbers to workers
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		startNum := i*chunkSize + 1
		endNum := startNum + chunkSize - 1

		// Ensure the last worker handles any remainder
		if i == numOfWorkers-1 {
			endNum = n
		}
		//wg.Add(1)
		go printNums(startNum, endNum, &wg)
	}

	wg.Wait()
	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}
