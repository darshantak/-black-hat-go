package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func processData(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done()
	processedData := data * 2

	lock.Lock()
	*result = append(*result, processedData)
	lock.Unlock()
}

func runMutex() {
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}

	//the result array is the shared resource
	result := []int{}
	//TODO: put the input into result

	for _, i := range input {
		wg.Add(1)
		go processData(&wg, &result, i)
	}

	wg.Wait()
	fmt.Println(result)
}
