package main

import (
	"fmt"
	"sync"
	"time"
)

func processWithTime(n int) int {
	// time.Sleep(time.Second * 2)
	return n * 2
}

func processData(wg *sync.WaitGroup, result *int, data int) {
	defer wg.Done()
	processedData := processWithTime(data)

	*result = processedData
}

func runMutex() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}

	//the result array is the shared resource
	result := make([]int, len(input))
	//TODO: put the input into result

	for i, data := range input {
		wg.Add(1)
		go processData(&wg, &result[i], data)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
	fmt.Println(result)
}
