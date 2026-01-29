package main

import "sync"

func Example() bool {
	wg := &sync.WaitGroup{}

	for x := 0; x <10; x++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			//
		}()
	}

	wg.Wait()

	return true
}