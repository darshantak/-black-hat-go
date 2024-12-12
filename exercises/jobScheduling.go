package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Job interface {
	Run()
}

type EmailJob struct {
	Message string
}

func (e EmailJob) Run() {
	fmt.Printf("Starting EmailJob: %s\n", e.Message)
	time.Sleep(1 * time.Second)
	fmt.Printf("Finished EmailJob: %s\n", e.Message)
}

type DataJob struct {
	Data string
}

func (d DataJob) Run() {
	fmt.Printf("Starting DataProcessingJob: %s\n", d.Data)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished DataProcessingJob: %s\n", d.Data)
}

type ReportJob struct {
	Report string
}

func (r ReportJob) Run() {
	fmt.Printf("Starting ReportJob: %s\n", r.Report)
	time.Sleep(3 * time.Second)
	fmt.Printf("Finished ReportJob: %s\n", r.Report)
}

func JobScheduler(jobMap map[int][]Job) {
	var priorities []int
	var wg sync.WaitGroup
	for priority := range jobMap {
		priorities = append(priorities, priority)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(priorities)))

	jobChan := make(chan Job)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobChan {
				job.Run()
			}
		}()
	}

	go func() {
		for _, prior := range priorities {
			for _, job := range jobMap[prior] {
				jobChan <- job
			}
		}
	}()
}
