package main

func Example(stopCh chan struct{}) {
	semaphore := make(chan struct{})

	for {
		select {
		case semaphore <- struct{}{}:
			go func ()  {
				doWork()

				<- semaphore
			}()
			case <- stopCh:
				return
		}
	}
}

func doWork() {}