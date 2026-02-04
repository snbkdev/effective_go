package main

import "sync"

func Example() {
	outputCh := make(chan Data)
	mutex := &sync.Mutex{}

	mutex.Lock()

	outputCh <- Data{}

	mutex.Unlock()
}

type Data struct {}