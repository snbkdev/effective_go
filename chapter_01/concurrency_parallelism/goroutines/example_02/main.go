package main

import (
	"runtime"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	semaphore := make(chan struct{}, runtime.NumCPU())

	for x := 0; x < 100000; x++ {
		wg.Add(1)
		go doWork(wg, semaphore)
	}

	wg.Wait()
}

func doWork(wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()

	semaphore <- struct{}{}

	sum := 0

	for x := 0; x < 1000000; x++ {
		sum += x

		if x % 100 == 0 {
			runtime.Gosched()
		}
	}

	<-semaphore
}