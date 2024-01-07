package main

import (
	"log"
	"sync"
	"time"
)

var (
	mutex  sync.Mutex      = sync.Mutex{}
	result []int           = []int{}
	wg     *sync.WaitGroup = &sync.WaitGroup{}
)

// print numbers
// print result: 3, 2, 1
// defer functions are executed in reverse order
func printNo() {
	defer log.Println("printNo=1")
	defer log.Println("printNo=2")
	defer log.Println("printNo=3")

	log.Println("run defer functions after 3 seconds")

	time.Sleep(time.Duration(3) * time.Second)
}

// print numbers
// print result: 2, 3, 1
// defer functions are executed in reverse order
func printNo2() {
	defer log.Println("printNo=1")
	defer log.Println("printNo=3")
	defer log.Println("printNo=2")

	log.Println("run defer functions after 3 seconds")

	time.Sleep(time.Duration(3) * time.Second)
}

// raise panic
func raisePanic() {
	log.Println("raise panic after 3 seconds")
	time.Sleep(time.Duration(3) * time.Second)

	panic("raisePanic!")
}

// recover panic
func recoverPanic() {
	defer func() {
		log.Println("post panic event")
	}()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic occurred, error=%v", r)
		}
	}()
	defer func() {
		log.Println("pre panic event")
	}()

	raisePanic()
}

// defer goroutine
// print result: random
func goroutineDefer() {
	// use i in for-loop may not properly work on closure with goroutine due to bug in golang before 1.22
	// this bug will be fixed in 1.22 (ref: https://go.dev/blog/loopvar-preview, https://go.dev/doc/faq#closures_and_goroutines)
	for i := 0; i <= 4; i++ {
		j := i
		defer func() {
			go log.Printf("goroutine defer exit=%d\n", j)
		}()
	}

	log.Println("run goroutine in defer functions after 3 seconds")
	time.Sleep(time.Duration(3) * time.Second)
}

func main() {
	log.Println("run defer order")

	// print number test
	printNo()
	printNo2()

	// raise recover
	recoverPanic()

	// sleep 3 seconds to ensure goroutines are executed
	goroutineDefer()
	time.Sleep(time.Duration(1) * time.Second)
}
