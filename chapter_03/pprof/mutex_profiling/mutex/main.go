package main

import "sync"

var (
	counter int
	mutex = &sync.Mutex{}
)

func main() {
	wg := &sync.WaitGroup{}

	for x := 0; x < 1000; x++ {
		wg.Add(1)
		go performTask(wg)
	}

	wg.Wait()
}

func performTask(wg *sync.WaitGroup) {
	defer wg.Done()

	for x := 0; x < 100000; x++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}