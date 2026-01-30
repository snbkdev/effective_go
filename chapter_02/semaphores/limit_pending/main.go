package main

import "sync"

func Example(stopCh chan struct{}, workCh chan Data) {
	wg := &sync.WaitGroup{}

	semaphore := make(chan struct{}, 10)

	for {
		select {
		case data := <- workCh:
			wg.Add(1)

			semaphore <- struct{}{}

			go func() {
				defer wg.Done()

				doWork(data)

				<- semaphore
			}()

			case <- stopCh:
				wg.Wait()

				return
		}
	}
}

type Data struct{}

func doWork(data Data) {}