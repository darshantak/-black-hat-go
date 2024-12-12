package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func readUrlFile() []string {
	file, err := os.Open("./tempFiles/urls.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	var urlList []string
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		urlList = append(urlList, line)
	}
	return urlList
}

func scarpeHelper(url string, wg *sync.WaitGroup) bool {
	defer wg.Done()
	fmt.Println(url)
	return true
}

func scrapeURL() {
	var wg sync.WaitGroup

	urlList := readUrlFile()
	for _, url := range urlList {
		wg.Add(1)
		go scarpeHelper(url, &wg)
	}
	wg.Wait()
}
