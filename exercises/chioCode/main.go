package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func workerWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	/**
	* Goroutines Patterns
	**/

	// Channels
	done := make(chan bool, 1)
	go worker(done)

	<-done

	fmt.Println()
	fmt.Println()

	// Using SELECT
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received: ", msg1)
		case msg2 := <-c2:
			fmt.Println("Received: ", msg2)
		}
	}

	fmt.Println()
	fmt.Println()

	wellDone := make(chan bool)
	go func() {
		fmt.Println("Wooooorking...!")
		time.Sleep(3 * time.Second)
		wellDone <- true
	}()

	select {
	case <-wellDone:
		fmt.Println("Work ended!")
	case <-time.After(2 * time.Second):
		fmt.Println("Time out: more two seconds!")
	}

	fmt.Println()
	fmt.Println()

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerWaitGroup(i, &wg)
	}

	wg.Wait()
}
