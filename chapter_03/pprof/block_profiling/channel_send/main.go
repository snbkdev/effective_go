package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.SetBlockProfileRate(1)

	semaphore := make(chan struct{}, 5)

	wg := &sync.WaitGroup{}

	for x := 0; x < 100; x++ {
		wg.Add(1)
		go performTask(wg, semaphore)
	}

	wg.Wait()

	fmt.Println("All done!")

	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
}

func performTask(wg *sync.WaitGroup, semaphore chan struct{}) {
	defer func() {
		<-semaphore

		wg.Done()
	}()

	semaphore <- struct{}{}

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