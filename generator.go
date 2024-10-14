package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func infiniteGenerator[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func primeFinder(done <-chan int, randIntSteam <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randInt := <-randIntSteam:
				if isPrime(randInt) {
					primes <- randInt
				}
			}
		}
	}()
	return primes
}

func fanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInSteam := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInSteam <- i:
			}
		}
	}
	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}
	go func() {
		wg.Wait()
		close(fannedInSteam)
	}()

	return fannedInSteam
}
func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()
	return taken
}

func runGenerator() {
	start := time.Now()
	done := make(chan int)
	defer close(done)
	i := 1
	randomNumGenerator := func() int {
		val := i
		i++
		return val
		// return rand.Intn(500)
	}

	numStream := infiniteGenerator(done, randomNumGenerator)

	// primeNumStream := primeFinder(done, numStream)

	//fan out we have to limit the number of the CPUs for this to scale up to how many routines.
	CPUcount := runtime.NumCPU()
	//fan out channel
	primeFinderChannel := make([]<-chan int, CPUcount)

	for i := 0; i < CPUcount; i++ {
		primeFinderChannel[i] = primeFinder(done, numStream)
	}

	//Fan in the channels
	fannedStream := fanIn(done, primeFinderChannel...)

	for n := range take(done,fannedStream,10) {
		fmt.Println(n)
	}
	// Naive approach
	// for n := range take(done, primeNumStream, 10) {
	// 	fmt.Println(n)
	// }
	end := time.Since(start)
	fmt.Println(end)
}
