package main

import (
	"runtime"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	for x := 0; x < 100000; x++ {
		wg.Add(1)
		go doWorking(wg)
	}

	wg.Wait()
}

func doWorking(wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0

	for x := 0; x < 100000; x++ {
		sum += x

		if x % 100 == 0 {
			runtime.Gosched()
		}
	}
}