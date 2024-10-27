// Exercise 6: Fan-In and Fan-Out
// Implement a fan-in/fan-out pattern:
// Create two producer goroutines that generate numbers (one produces even numbers and the other produces odd numbers).
// Use a fan-in function to merge the output from these producers into a single channel.
// Create multiple worker goroutines (fan-out) that consume the numbers and print them.

// Exercise 7: Bounded Parallelism
// Write a program that limits the number of concurrently running goroutines. For example, the program should process a list of URLs and only allow 3 goroutines to run at a time. Use channels or semaphores to manage this constraint.

// Exercise 8: Deadlock Detection and Resolution
// Write two goroutines that each try to acquire two locks in different orders (potential deadlock). Identify the deadlock situation, and modify your program to resolve the deadlock using proper locking mechanisms or tryLock.
package exercises

import (
	"fmt"
	"sync"
)

type Number struct {
	Value  int
	Source string
}

func genEven(done <-chan int) <-chan Number {
	// i := 0
	// isEven := func() int {
	// 	for {
	// 		if i%2 == 0 {
	// 			val := i
	// 			i += 1
	// 			return val
	// 		}
	// 		i += 1
	// 	}
	// }
	evenStream := make(chan Number)
	go func() {
		defer close(evenStream)
		i := 0
		for i < 10 {
			select {
			case <-done:
				return
				// case evenStream <- isEven():
			case evenStream <- Number{Value: i, Source: "Even"}:
				i += 2
			}
		}
	}()
	return evenStream
}
func genOdd(done <-chan int) <-chan Number {
	// i := 0
	// isOdd := func() int {
	// 	for {
	// 		if i%2 != 0 {
	// 			val := i
	// 			i += 1
	// 			return val
	// 		}
	// 		i += 1
	// 	}
	// }
	oddStream := make(chan Number)
	go func() {
		defer close(oddStream)
		i := 1
		for i < 10 {
			select {
			case <-done:
				return
			// case oddStream <- isOdd():
			case oddStream <- Number{Value: i, Source: "Odd"}:
				i += 2
			}
		}
	}()
	return oddStream
}
func fanInEvenOdd(done <-chan int, stream ...<-chan Number) <-chan Number {
	var wg sync.WaitGroup
	fannedInStream := make(chan Number)

	transfer := func(c <-chan Number) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}
	for _, c := range stream {
		wg.Add(1)
		go transfer(c)
	}
	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}
func RunExec6() {
	done := make(chan int)
	evenFanOut := make([]<-chan Number, 2)
	oddFanOut := make([]<-chan Number, 2)
	for i := 0; i < 2; i++ {
		evenFanOut[i] = genEven(done)
		oddFanOut[i] = genOdd(done)
	}

	oddStream := fanInEvenOdd(done, oddFanOut...)
	evenStream := fanInEvenOdd(done, evenFanOut...)

	finalStream := fanInEvenOdd(done, oddStream, evenStream)

	for c := range finalStream {
		fmt.Println(c)
	}
	close(done)
}
