package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func helper(path <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for filePath := range path {
		fmt.Printf("Processing file: %s\n", filePath)
		time.Sleep(1 * time.Second)
		fmt.Printf("Finished processing file: %s\n", filePath)
	}
}

func processFile() {
	filePathChannel := make(chan string)
	files, err := os.ReadDir("./tempFiles")

	if err != nil {
		log.Fatal(err.Error())
	}
	var wg sync.WaitGroup

	for j := 0; j < 3; j++ {
		wg.Add(1)
		go helper(filePathChannel, &wg)
	}

	go func() {
		for _, fileName := range files {
			filePathChannel <- fileName.Name()
		}
		close(filePathChannel)
	}()

	wg.Wait()
}
