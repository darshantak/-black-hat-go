package exercises

//Write a program that launches two goroutines.
//One goroutine should print numbers from 1 to 5 with a 100-millisecond delay between each number,
//and the other goroutine should print letters from 'a' to 'e' with a 200-millisecond delay.
//Both goroutines should run concurrently.

import (
	"fmt"
	"sync"
	"time"
)

func printNumbersRange(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < n+1; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := 'a'; ch <= 'e'; ch++ {
		fmt.Printf("%c   ", ch)
		time.Sleep(time.Millisecond * 200)
	}
}
func RunExer() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbersRange(5, &wg)
	go printLetters(&wg)
 
	wg.Wait()
}
