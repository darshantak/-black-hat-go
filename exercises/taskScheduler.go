package main

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Run()
}
type PrintTask struct {
	Message string
}

func (p PrintTask) Run() {
	fmt.Printf("Starting PrintTask: %s\n", p.Message)
	time.Sleep(1 * time.Second)
	fmt.Printf("Finished PrintTask: %s\n", p.Message)
}

type DelayTask struct {
	Duration time.Duration
}

func (d DelayTask) Run() {
	fmt.Printf("Starting DelayTask for %d seconds\n", d.Duration)
	time.Sleep(d.Duration * time.Second)
	fmt.Printf("Finished DelayTask for %d seconds\n", d.Duration)
}

func TaskScheduler(tasks []Task) {
	var wg sync.WaitGroup

	taskChannel := make(chan Task)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskChannel {
				task.Run()
			}
		}()
	}

	go func() {
		for _, i := range tasks {
			taskChannel <- i
		}
		close(taskChannel)
	}()
	wg.Wait()
}
func RunMain() {
	tasks := []Task{
		PrintTask{Message: "Hey"},
		DelayTask{Duration: 10},
	}
	TaskScheduler(tasks)
}
