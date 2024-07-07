package main

import (
	"fmt"
	"time"
)

func routine() {
	count := 0
	for i := 0; i < 1000; i++ {
		count += 1
	}

	fmt.Println("This is the first go routine")
}

func main() {
	go routine()
	fmt.Println("Hi, this is the playground for Go")
	time.Sleep(1000 * time.Millisecond)
}
