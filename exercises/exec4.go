package exercises

import (
	"fmt"
	"sync"
)

func incrementCounter(counter *int, wg *sync.WaitGroup) {
	defer wg.Done()
	// lock.Lock()
	*counter += 1
	// lock.Unlock()
}

func stage1(n int) <-chan int {
	numChan := make(chan int)
	go func() {
		for i := 1; i < n+1; i++ {
			numChan <- i
		}
		close(numChan)
	}()
	return numChan
}

func stage2(numStream <-chan int) <-chan int {
	multiChan := make(chan int)
	go func() {
		for i := range numStream {
			multiChan <- i * 2
		}
		close(multiChan)
	}()
	return multiChan
}

func RunExec5() {
	//Stage 1 Generate number channel from 1-5
	integerStreamChan := stage1(5)
	//Stage 2 Multiply by 2 for every integer in integerStreamChan
	multipliedStreamChan := stage2(integerStreamChan)
	//Stage 3 Printing all the numbers in the multipliedStreamChan
	for i := range multipliedStreamChan {
		fmt.Println(i)
	}
}
func RunExec4() {
	// Write a program where multiple goroutines increment a shared counter.
	// Use a sync.Mutex to ensure that the counter is incremented safely without race conditions.
	// The main function should wait for all goroutines to complete and then print the final value of the counter.
	var wg sync.WaitGroup
	counter := 0
	wg.Add(12)
	go incrementCounter(&counter, &wg)
	go incrementCounter(&counter, &wg)
	go incrementCounter(&counter, &wg)
	go incrementCounter(&counter, &wg)
	go incrementCounter(&counter, &wg)
	go incrementCounter(&counter, &wg)
	go incrementCounter(&counter, &wg)
	go incrementCounter(&counter, &wg)
	go incrementCounter(&counter, &wg)
	go incrementCounter(&counter, &wg)
	go incrementCounter(&counter, &wg)
	go incrementCounter(&counter, &wg)
	wg.Wait()
	fmt.Println(counter)
}
