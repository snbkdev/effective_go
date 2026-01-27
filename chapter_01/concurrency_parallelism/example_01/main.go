package main

import (
	"runtime"
	"sync"
)

func output(wg *sync.WaitGroup, value int, result chan int) {
	defer wg.Done()

	for x := 0; x < 15; x++ {
		result <- value

		runtime.Gosched()
	}
}

func main() {
	result := make(chan int, 100)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go output(wg, 0, result)

	wg.Add(1)
	go output(wg, 1, result)

	wg.Wait()

	close(result)

	for value := range result {
		print(value)
	}
}