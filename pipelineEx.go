package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _,n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func runPipeline() {
	nums := []int{1, 2, 3, 4, 5}
	//Stage 1
	dataChannel := sliceToChannel(nums)

	//Stage 2
	squared := sq(dataChannel)

	//Stage 3
	for i := range squared {
		fmt.Println(i)
	}
}
