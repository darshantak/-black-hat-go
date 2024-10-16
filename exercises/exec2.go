package exercises

import (
	"fmt"
	"sync"
)

//Create a program where two goroutines communicate via a channel. One goroutine should send a series of numbers to a channel,
//and the other should receive those numbers from the channel and print them out.


func RunExec2() {
	numberChannel := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			numberChannel <- i
		}
		close(numberChannel)
	}()

	go func() {
		wg.Done()
		for i := range numberChannel {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}
