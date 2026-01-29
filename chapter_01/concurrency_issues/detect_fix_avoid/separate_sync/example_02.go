package main

func UpdateDestinationV2() {
	sourceMutex.Lock()
	results := doCalculation(sourceMap)
	sourceMutex.Unlock()

	destinationMutex.Lock()
	destinationMap = results
	destinationMutex.Unlock()
}