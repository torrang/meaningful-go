package main

import (
	"log"
	"sync"
	"time"
)

// user defined asynchronous task
func async_task(wg *sync.WaitGroup, stopEvent <-chan bool, i int) {
	for {
		select {
		case <-stopEvent:
			// using break does not escape for-loop
			// should use return instead or flag
			log.Printf("terminate async task %d", i)
			wg.Done()
			return
		default:
			log.Printf("async task %d is running", i)
			time.Sleep(time.Duration(i) * time.Second)
		}
	}
}

func main() {
	log.Printf("start channel-goroutine-manage")
	waitGroup := sync.WaitGroup{}
	stopEvent := make(chan bool)

	// run goroutines
	for i := 1; i <= 3; i++ {
		waitGroup.Add(1)
		go async_task(&waitGroup, stopEvent, i)
	}

	// cancel goroutines after 10 seconds
	go func() {
		log.Printf("wait for 10 seconds")
		time.Sleep(time.Duration(10) * time.Second)
		log.Printf("time is over, close channel")
		close(stopEvent)
	}()

	// wait goroutines
	log.Printf("wait goroutines")
	waitGroup.Wait()

	log.Printf("all goroutines are terminated")
	log.Printf("terminate basic-goroutine-manage")
}
