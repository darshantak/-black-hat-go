package exercises

import (
	"fmt"
	"sync"
	"time"
)

//Write a program that creates a worker pool.
//There should be a pool of 5 workers, and the main function should send 10 tasks to the workers via a channel.
//Each task should simulate a job by sleeping for 500 milliseconds and then printing "Task done".

func workerFunc(taskChannel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range taskChannel {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(i, "Task Done")
	}
}
func RunExec3() {
	var wg sync.WaitGroup
	taskChannel := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go workerFunc(taskChannel, &wg)
	}

	for i := 0; i < 10; i++ {
		taskChannel <- i
	}
	close(taskChannel)

	wg.Wait()
}
