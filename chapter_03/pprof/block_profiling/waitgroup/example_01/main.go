package main

import (
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	for x := 0; x < 10; x++ {
		wg.Add(1)

		go performTask(wg)
	}

	wg.Wait()
}

func performTask(wg *sync.WaitGroup) {
	defer wg.Done()

	done := time.After(1 * time.Second)
	x := 0

	for {
		x++

		select {
		case <-done:
			return

		default:
			x++
		}
	}
}