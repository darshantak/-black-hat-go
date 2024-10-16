package main

import "sync"

var wg sync.WaitGroup

func runDone() {

	done := make(chan interface{})
	defer close(done)

	cows := make(chan interface{}, 100)
	// pigs := make(chan interface{}, 100)

	go func() {
		for {
			cows <- "moo"
		}
	}()

	go func() {
		for {
			cows <- "oink"
		}
	}()

	wg.Add(1)
	// go func() {
	// 	done <- "string"
	// 	done <- 24
	// }()
	// var1 := <-done
	// var2 := <-done
	// fmt.Println(var1)
	// fmt.Println(var2)
	// time.Sleep(time.Second*2)
}
