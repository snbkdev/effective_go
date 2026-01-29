package main

import (
	"fmt"
	"sync"
	"time"
)

var inOnline = true
var isOnlineMutex = &sync.Mutex{}

func connection(name string, serverStatus *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := 0; x < 30; x++ {
		serverStatus.L.Lock()

		for !inOnline {
			fmt.Printf("%s: blocked\n", name)
			serverStatus.Wait()
		}

		serverStatus.L.Unlock()

		sendToServer(name, x)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	serverStatus := sync.NewCond(isOnlineMutex)

	wg.Add(1)
	go connection("A", serverStatus, wg)

	wg.Add(1)
	go connection("B", serverStatus, wg)

	<- time.After(1 * time.Second)

	serverStatus.L.Lock()
	inOnline = false
	serverStatus.L.Unlock()

	<- time.After(1 * time.Second)

	serverStatus.L.Lock()
	inOnline = true
	serverStatus.L.Unlock()

	serverStatus.Broadcast()

	wg.Wait()
}

func sendToServer(name string, index int) {
	fmt.Printf("%s: %d\n", name, index)

	<- time.After(100 * time.Millisecond)
}