package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.SetBlockProfileRate(1)

	dataCh := make(chan int)
	go publishData(dataCh)

	wg := &sync.WaitGroup{}
	for x := 0; x < 100; x++ {
		wg.Add(1)

		go consumeData(wg, dataCh)
	}

	wg.Wait()

	fmt.Println("All done!")

	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
}

func publishData(dataCh chan int) {
	for x := 0; x < 1000; x++ {
		dataCh <- rand.Int()

		time.Sleep(10 * time.Millisecond)
	}

	close(dataCh)
}

func consumeData(wg *sync.WaitGroup, dataCh chan int) {
	defer wg.Done()

	for value := range dataCh {
		fmt.Printf("Value: %d\n", value)
	}
}