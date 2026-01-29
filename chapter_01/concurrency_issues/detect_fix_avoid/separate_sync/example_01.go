package main

import "sync"

var (
	sourceMap   = map[string]string{}
	sourceMutex = &sync.Mutex{}

	destinationMap   = map[string]string{}
	destinationMutex = &sync.Mutex{}
)

func UpdateDestinationV1() {
	sourceMutex.Lock()
	destinationMutex.Lock()

	destinationMap = doCalculation(sourceMap)

	destinationMutex.Unlock()
	sourceMutex.Unlock()
}

func doCalculation(in map[string]string) map[string]string {
	return in
}