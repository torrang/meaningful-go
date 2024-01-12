package main

import (
	"log"
	"time"
)

func main() {
	log.Println("closure")

	// closure in for loop work well
	// as you can see, address of i is same
	// meaning that integer 0 to 10 is assigned to same variable not creating new variable
	for i := 0; i <= 10; i++ {
		func() {
			log.Printf("i=%d, &i=%p\n", i, &i)
		}()
	}

	log.Println("closure with goroutines")

	// closure with goroutine in for loop not work well
	// still assigned to same variable
	// for loop is finished before any of goroutines start
	// make it print unpredictable value between 1 to 10
	for i := 0; i < 10; i++ {
		go func() {
			log.Printf("i=%d, &i=%p\n", i, &i)
		}()
	}

	time.Sleep(1 * time.Second)

}
