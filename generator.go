package main

import (
	"fmt"
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
func runGenerator() {
	done := make(chan bool)
	i := 1
	randomNumGenerator := func() int {
		val := i
		i++
		return val
		// return rand.Intn(500)
	}

	numStream := infiniteGenerator(done, randomNumGenerator)

	for n := range numStream {
		fmt.Println(n)
		if n == 5 {
			close(done)
		}
	}
}
